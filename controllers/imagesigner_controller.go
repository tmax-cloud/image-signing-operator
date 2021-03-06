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
	"github.com/tmax-cloud/image-signing-operator/pkg/controller"
	"github.com/tmax-cloud/image-signing-operator/pkg/trust"
)

// ImageSignerReconciler reconciles a ImageSigner object
type ImageSignerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tmax.io,resources=imagesigners,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tmax.io,resources=imagesigners/status,verbs=get;update;patch

func (r *ImageSignerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("imagesigner", req.NamespacedName)

	// get image signer
	signer := &tmaxiov1.ImageSigner{}
	if err := r.Get(context.TODO(), req.NamespacedName, signer); err != nil {
		log.Error(err, "")
		return ctrl.Result{}, nil
	}

	if signer.Status.SignerKeyState != nil && signer.Status.Created {
		return ctrl.Result{}, nil
	}

	defer updateSignerStatus(r.Client, signer)

	// check if signer key is exist
	signerKey := &tmaxiov1.SignerKey{}
	r.Get(context.TODO(), types.NamespacedName{Name: signer.Name}, signerKey)
	if len(signerKey.Name) > 0 {
		log.Info("signer key is already exist")
		return ctrl.Result{}, nil
	}

	// if signer key is not exist, create root key
	signCtl := controller.NewSigningController(r.Client, signer, "", "", "")
	phrase := trust.NewTrustPass()
	phrase.AssignNewRootPass()
	cmdOpt := &controller.CommandOpt{
		RootKey: &tmaxiov1.TrustKey{PassPhrase: phrase[trust.DctEnvKeyRoot]},
	}
	log.Info("dind start")
	if err := signCtl.Start(cmdOpt); err != nil {
		log.Error(err, "dind container start failed")
		makeSignerStatus(signer, false, err.Error(), "", nil)
		signCtl.Close()
		return ctrl.Result{}, nil
	}
	defer signCtl.Close()

	if !signCtl.IsRunnging {
		makeSignerStatus(signer, false, "dind pod is not running", "", nil)
		return ctrl.Result{}, nil
	}
	log.Info("dind is running")

	rootKey, err := signCtl.CreateRootKey(phrase, signer, r.Scheme)
	if err != nil {
		makeSignerStatus(signer, false, err.Error(), "", nil)
		return ctrl.Result{}, nil
	}

	makeSignerStatus(signer, true, "", "", rootKey)
	if signer.Status.SignerKeyState == nil {
		log.Info("SignerKeyState is nil!!!!")
	}

	return ctrl.Result{}, nil
}

func (r *ImageSignerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tmaxiov1.ImageSigner{}).
		Owns(&tmaxiov1.SignerKey{}).
		Complete(r)
}
