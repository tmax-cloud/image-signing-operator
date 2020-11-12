/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tmaxiov1 "github.com/tmax-cloud/image-signing-operator/api/v1"
	"github.com/tmax-cloud/image-signing-operator/internal/utils"
	"github.com/tmax-cloud/image-signing-operator/pkg/controller"
	"github.com/tmax-cloud/image-signing-operator/pkg/trust"
)

// ImageSignRequestReconciler reconciles a ImageSignRequest object
type ImageSignRequestReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tmax.io,resources=imagesignrequests,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tmax.io,resources=imagesignrequests/status,verbs=get;update;patch

func (r *ImageSignRequestReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("imagesignrequest", req.NamespacedName)

	// get image sign request
	log.Info("get image sign request")
	signReq := &tmaxiov1.ImageSignRequest{}
	if err := r.Get(context.TODO(), req.NamespacedName, signReq); err != nil {
		log.Error(err, "")
		makeResponse(signReq, false, err.Error(), "")
		return ctrl.Result{}, nil
	}

	defer response(r.Client, signReq)

	// get image signer
	log.Info("get image signer")
	signer := &tmaxiov1.ImageSigner{}
	if err := r.Get(context.TODO(), types.NamespacedName{Name: signReq.Spec.Signer}, signer); err != nil {
		log.Error(err, "")
		makeResponse(signReq, false, err.Error(), "")
		return ctrl.Result{}, nil
	}

	// get sign key
	log.Info("get sign key")
	signerKey := &tmaxiov1.SignerKey{}
	if err := r.Get(context.TODO(), types.NamespacedName{Name: signReq.Spec.Signer}, signerKey); err != nil {
		log.Error(err, "")
		makeResponse(signReq, false, err.Error(), "")
		return ctrl.Result{}, nil
	}

	// get trust pass
	log.Info("get trust pass")
	phrase := make(trust.TrustPass)
	addedTargetKey := false
	phrase[trust.DctEnvKeyRoot] = signerKey.Spec.Root.PassPhrase
	if _, ok := signerKey.Spec.Targets[buildTargetName(signReq)]; ok {
		phrase[trust.DctEnvKeyTarget] = signerKey.Spec.Targets[signReq.Spec.Image].PassPhrase
	} else {
		phrase.AssignNewTargetPass()
		addedTargetKey = true
	}

	//
	signCtl := controller.NewSigningController(r.Client, signer, phrase, signReq.Spec.RegistryLogin.Name, signReq.Spec.RegistryLogin.Namespace, req.Namespace)
	cmdOpt := &controller.CommandOpt{
		Phrase:                  signCtl.Phrase,
		RegistryLoginSecret:     signReq.Spec.RegistryLogin.DcjSecretName,
		RegistryLoginCertSecret: signReq.Spec.RegistryLogin.CertSecretName,
		ImagePvc:                signReq.Spec.PvcName,
	}

	log.Info("dind start")
	if err := signCtl.Start(cmdOpt); err != nil {
		log.Error(err, "dind container start failed")
		makeResponse(signReq, false, err.Error(), "")
		return ctrl.Result{}, nil
	}
	defer signCtl.Close()

	if !signCtl.IsRunnging {
		return ctrl.Result{}, nil
	}
	log.Info("dind is running")

	log.Info("sign image")
	imageName, imageTag := utils.ParseImage(signReq.Spec.Image)
	if err := signCtl.SignImage(imageName, imageTag); err != nil {
		makeResponse(signReq, false, err.Error(), "")
		return ctrl.Result{}, nil
	}

	if addedTargetKey {
		log.Info("add target key to signerkey")
		if err := signCtl.AddTargetKey(
			signerKey,
			buildTargetName(signReq),
			phrase,
		); err != nil {
			makeResponse(signReq, false, err.Error(), "")
			return ctrl.Result{}, nil
		}
	}

	makeResponse(signReq, true, "", "")
	return ctrl.Result{}, nil
}

func (r *ImageSignRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tmaxiov1.ImageSignRequest{}).
		Complete(r)
}

func buildTargetName(signReq *tmaxiov1.ImageSignRequest) string {
	imageName, _ := utils.ParseImage(signReq.Spec.Image)
	return trust.BuildTargetName(
		signReq.Spec.RegistryLogin.Name,
		signReq.Spec.RegistryLogin.Namespace,
		imageName,
	)
}
