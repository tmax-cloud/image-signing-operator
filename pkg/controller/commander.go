package controller

import (
	"bytes"
	"os"
	"path"
	"strings"

	"github.com/tmax-cloud/image-signing-operator/internal/k8s"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ExecResult is execute result buffer
type ExecResult struct {
	Outbuf *bytes.Buffer
	Errbuf *bytes.Buffer
}

const (
	// BaseDir is docker trust content base directory path
	BaseDir = "/root/.docker/trust"
	// PrivateKeyDir is docker trust content private key directory path
	PrivateKeyDir = BaseDir + "/private"
)

// KubeCommander is a commander to excute command to container in specified pod
type KubeCommander struct {
	client                    client.Client
	namespace, pod, container string
}

// NewKubeCommander is
func NewKubeCommander(c client.Client, namespace, pod string) *KubeCommander {
	if len(namespace) == 0 {
		namespace = os.Getenv("OPERATOR_NAMESPACE")
	}

	return &KubeCommander{
		client:    c,
		namespace: namespace,
		pod:       pod,
		container: "docker-cli",
	}
}

// GenerateKey generates new key
func (k *KubeCommander) GenerateKey(role string) (*ExecResult, error) {
	command := []string{"mkdir", "-p", BaseDir}
	if _, err := k.excute(strings.Join(command, " ")); err != nil {
		return nil, err
	}

	command = []string{"docker", "trust", "key", "generate", role, "--dir", BaseDir}
	return k.excute(strings.Join(command, " "))
}

// ListKey returns key list in /root/.docker/trust/private directory
func (k *KubeCommander) ListKey() (*ExecResult, error) {
	command := []string{"ls", "--color=never", PrivateKeyDir}
	return k.excute(strings.Join(command, " "))
}

// ReadKey returns file content in /root/.docker/trust/private directory
func (k *KubeCommander) ReadKey(name string) (*ExecResult, error) {
	command := []string{"cat", path.Join(PrivateKeyDir, name)}
	return k.excute(strings.Join(command, " "))
}

// LoadImageTar loads tar image
func (k *KubeCommander) LoadImageTar(path string) (*ExecResult, error) {
	command := []string{"docker", "load", "<", path}
	return k.excute(strings.Join(command, " "))
}

// TagImage tags image from "target" to "tagName"
func (k *KubeCommander) TagImage(target, tagName string) (*ExecResult, error) {
	command := []string{"docker", "tag", target, tagName}
	return k.excute(strings.Join(command, " "))
}

// Sign executes sign and push image
func (k *KubeCommander) Sign(imageName string) (*ExecResult, error) {
	command := []string{"docker", "trust", "sign", imageName}
	return k.excute(strings.Join(command, " "))
}

// ListImageId is
func (k *KubeCommander) ListImageId() (*ExecResult, error) {
	command := []string{"docker", "images", "-q"}
	return k.excute(strings.Join(command, " "))
}

func (k *KubeCommander) excute(command string) (*ExecResult, error) {
	res := &ExecResult{Outbuf: &bytes.Buffer{}, Errbuf: &bytes.Buffer{}}
	if err := k8s.ExecCmd(k.pod, k.container, k.namespace, command, nil, res.Outbuf, res.Errbuf); err != nil {
		return nil, err
	}

	return res, nil
}
