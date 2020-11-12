package controller

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/go-logr/logr"
	apiv1 "github.com/tmax-cloud/image-signing-operator/api/v1"
	"github.com/tmax-cloud/image-signing-operator/internal/schemes"
	"github.com/tmax-cloud/image-signing-operator/internal/utils"
	"github.com/tmax-cloud/image-signing-operator/pkg/registry"
	"github.com/tmax-cloud/image-signing-operator/pkg/trust"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

var log logr.Logger = ctrl.Log.WithName("signing-controller")

type CommandOpt struct {
	RootKey                                      *apiv1.TrustKey
	TargetKey                                    *apiv1.TrustKey
	RegistryLoginSecret, RegistryLoginCertSecret string
	ImagePvc                                     string
}

// NewSigningController is a controller for image signing.
// if registryName or registryNamespace is empty string, RegCtl is nil
// if requestNamespace is empty string, get operator's namepsace
func NewSigningController(c client.Client, signer *apiv1.ImageSigner, registryName, registryNamespace, requestNamespace string) *SigningController {
	return &SigningController{
		ImageSigner: signer,
		Cmder:       NewKubeCommander(c, requestNamespace, "image-signing-by-"+signer.Name+"-"+utils.RandomString(10)),
		Regctl:      registry.NewRegCtl(c, registryName, registryNamespace),
	}
}

type SigningController struct {
	ImageSigner *apiv1.ImageSigner
	Cmder       *KubeCommander
	Regctl      *registry.RegCtl
	startedPod  *corev1.Pod
	IsRunnging  bool
}

func storeFileShellCommand(filename, contents string) string {
	cmd := []string{"echo", "\"" + contents + "\"", ">", path.Join(PrivateKeyDir, filename)}

	return strings.Join(cmd, " ")
}

func (c *SigningController) Start(cmdOpt *CommandOpt) error {
	lifeCycleCmds := []string{}
	envs := map[string]string{}

	addEnvAndCmd := func(trustKey *apiv1.TrustKey, roleName trust.RoleType) {
		if trustKey != nil {
			if len(trustKey.PassPhrase) > 0 {
				envs[trust.RoleMap[roleName]] = trustKey.PassPhrase
			}
			if len(trustKey.ID) > 0 && len(trustKey.Key) > 0 {
				lifeCycleCmds = append(lifeCycleCmds, storeFileShellCommand(trustKey.ID, trustKey.Key))
			}
		}
	}

	addEnvAndCmd(cmdOpt.RootKey, trust.TrustRoleRoot)
	addEnvAndCmd(cmdOpt.TargetKey, trust.TrustRoleTarget)
	if len(cmdOpt.RegistryLoginSecret) > 0 {
		lifeCycleCmds = append(lifeCycleCmds, "cp /home/dockremap/.dockerconfigjson /root/.docker/config.json")
	}

	c.startedPod = schemes.NewDindPod(
		c.Cmder.namespace,
		c.Cmder.pod,
		c.Cmder.container,
		"",
		schemes.WithEnv(envs),
		schemes.WithPvc(cmdOpt.ImagePvc),
		schemes.WithDcjSecret(cmdOpt.RegistryLoginSecret),
		schemes.WithCertSecret(cmdOpt.RegistryLoginCertSecret),
		schemes.WithLifeCycle(lifeCycleCmds),
	)

	if err := c.Cmder.client.Create(context.TODO(), c.startedPod); err != nil {
		if errors.IsAlreadyExists(err) {
			return nil
		}
		return err
	}

	const MaxRetryCount = 60
	for cnt := 0; cnt < MaxRetryCount; cnt++ {
		if err := c.Cmder.client.Get(context.TODO(), client.ObjectKey{Name: c.Cmder.pod, Namespace: c.Cmder.namespace}, c.startedPod); err != nil {
			return err
		}
		if c.startedPod.Status.Phase == corev1.PodRunning {
			c.IsRunnging = true
			break
		}
		time.Sleep(1 * time.Second)
	}

	if !c.IsRunnging {
		return fmt.Errorf("pod is not running")
	}

	return nil
}

func (c *SigningController) Close() error {
	if err := c.Cmder.client.Delete(context.TODO(), c.startedPod); err != nil {
		return err
	}
	log.Info("dind closed", "pod/namespace", c.startedPod.Name+"/"+c.startedPod.Namespace)

	return nil
}

func (c *SigningController) readTrustKey(phrase trust.TrustPass, roleName trust.RoleType) (*apiv1.TrustKey, error) {
	log.Info("list key")
	out, err := c.Cmder.ListKey()
	if err != nil {
		log.Error(err, "list_key_err")
		return nil, err
	}
	log.Info("list key suceess", "stdout", out.Outbuf.String(), "stderr", out.Errbuf.String())

	keys := strings.Fields(out.Outbuf.String())
	keyFileFound := false
	trustKey := &apiv1.TrustKey{}

	for _, key := range keys {
		log.Info("private key", "key", key)
		readKeyOut, err := c.Cmder.ReadKey(key)
		if err != nil {
			log.Error(err, "")
			return nil, err
		}
		if strings.Contains(readKeyOut.Outbuf.String(), "role: "+string(roleName)) {
			trustKey.ID = key
			trustKey.Key = readKeyOut.Outbuf.String()
			trustKey.PassPhrase = phrase[trust.RoleMap[roleName]]
			keyFileFound = true
			break
		}
	}

	if !keyFileFound {
		return nil, fmt.Errorf("key file not found")
	}

	return trustKey, nil
}

func (c *SigningController) CreateRootKey(phrase trust.TrustPass, owner *apiv1.ImageSigner, scheme *runtime.Scheme) (*apiv1.TrustKey, error) {
	log.Info("generate key")
	out, err := c.Cmder.GenerateKey(string(trust.TrustRoleRoot))
	if err != nil {
		log.Error(err, "generate key err")
		return nil, err
	}
	log.Info("generate key success", "stdout", out.Outbuf.String(), "stderr", out.Errbuf.String())

	rootKey, err := c.readTrustKey(phrase, trust.TrustRoleRoot)
	if err != nil {
		log.Error(err, "read key err")
		return nil, err
	}

	log.Info("create root key")
	if err := c.createRootKey(owner, scheme, rootKey); err != nil {
		log.Error(err, "")
		return nil, err
	}

	log.Info("create root key success")
	return rootKey, nil
}

func (c *SigningController) AddTargetKey(originalKey *apiv1.SignerKey, targetName string, phrase trust.TrustPass) error {
	targetKey, err := c.readTrustKey(phrase, trust.TrustRoleTarget)
	if err != nil {
		log.Error(err, "read key error")
		return err
	}

	target := originalKey.DeepCopy()
	originObject := client.MergeFrom(originalKey)

	target.Spec.Targets[targetName] = *targetKey

	if err := c.Cmder.client.Patch(context.TODO(), target, originObject); err != nil {
		log.Error(err, "patch error")
		return err
	}

	return nil
}

func (c *SigningController) createRootKey(owner *apiv1.ImageSigner, scheme *runtime.Scheme, trustKey *apiv1.TrustKey) error {
	key := schemes.SignerKey(c.ImageSigner)
	if err := controllerutil.SetOwnerReference(owner, key, scheme); err != nil {
		return err
	}

	key.Spec = apiv1.SignerKeySpec{
		Root: *trustKey,
	}

	if err := c.Cmder.client.Create(context.TODO(), key); err != nil {
		return err
	}

	return nil
}

func (c *SigningController) SignImage(imageName, imageTag string) error {
	out, err := c.Cmder.LoadImageTar(path.Join(schemes.ImageMountPath, imageName+".tar"))
	if err != nil {
		log.Error(err, "load image error")
		return err
	}
	log.Info("load image", "stdout", out.Outbuf.String(), "stderr", out.Errbuf.String())

	out, err = c.Cmder.ListImageId()
	if err != nil {
		log.Error(err, "list image id error")
		return err
	}
	log.Info("list image id", "stdout", out.Outbuf.String(), "stderr", out.Errbuf.String())

	imageIds := strings.Fields(out.Outbuf.String())
	if len(imageIds) == 0 {
		return fmt.Errorf("image is not found")
	}

	registry := c.Regctl.GetEndpoint()

	image := path.Join(registry, imageName) + ":" + imageTag
	out, err = c.Cmder.TagImage(imageIds[0], image)
	if err != nil {
		log.Error(err, "list image id error")
		return err
	}
	log.Info("tag image", "stdout", out.Outbuf.String(), "stderr", out.Errbuf.String())

	log.Info("sign", "image name", image)
	out, err = c.Cmder.Sign(image)
	if err != nil {
		log.Error(err, "sign error")
		return err
	}
	log.Info("sign image", "stdout", out.Outbuf.String(), "stderr", out.Errbuf.String())

	return nil
}
