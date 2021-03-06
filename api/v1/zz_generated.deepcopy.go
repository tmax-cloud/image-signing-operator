// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/operator-framework/operator-lib/status"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreatePvc) DeepCopyInto(out *CreatePvc) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]AccessMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreatePvc.
func (in *CreatePvc) DeepCopy() *CreatePvc {
	if in == nil {
		return nil
	}
	out := new(CreatePvc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExistPvc) DeepCopyInto(out *ExistPvc) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExistPvc.
func (in *ExistPvc) DeepCopy() *ExistPvc {
	if in == nil {
		return nil
	}
	out := new(ExistPvc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignRequest) DeepCopyInto(out *ImageSignRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignRequest.
func (in *ImageSignRequest) DeepCopy() *ImageSignRequest {
	if in == nil {
		return nil
	}
	out := new(ImageSignRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSignRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignRequestList) DeepCopyInto(out *ImageSignRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageSignRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignRequestList.
func (in *ImageSignRequestList) DeepCopy() *ImageSignRequestList {
	if in == nil {
		return nil
	}
	out := new(ImageSignRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSignRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignRequestSpec) DeepCopyInto(out *ImageSignRequestSpec) {
	*out = *in
	out.RegistryLogin = in.RegistryLogin
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignRequestSpec.
func (in *ImageSignRequestSpec) DeepCopy() *ImageSignRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSignRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignRequestStatus) DeepCopyInto(out *ImageSignRequestStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignRequestStatus.
func (in *ImageSignRequestStatus) DeepCopy() *ImageSignRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ImageSignRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSigner) DeepCopyInto(out *ImageSigner) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSigner.
func (in *ImageSigner) DeepCopy() *ImageSigner {
	if in == nil {
		return nil
	}
	out := new(ImageSigner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSigner) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignerList) DeepCopyInto(out *ImageSignerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageSigner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignerList.
func (in *ImageSignerList) DeepCopy() *ImageSignerList {
	if in == nil {
		return nil
	}
	out := new(ImageSignerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSignerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignerSpec) DeepCopyInto(out *ImageSignerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignerSpec.
func (in *ImageSignerSpec) DeepCopy() *ImageSignerSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSignerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignerStatus) DeepCopyInto(out *ImageSignerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignerStatus.
func (in *ImageSignerStatus) DeepCopy() *ImageSignerStatus {
	if in == nil {
		return nil
	}
	out := new(ImageSignerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Registry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDeployment) DeepCopyInto(out *RegistryDeployment) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDeployment.
func (in *RegistryDeployment) DeepCopy() *RegistryDeployment {
	if in == nil {
		return nil
	}
	out := new(RegistryDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryList) DeepCopyInto(out *RegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryList.
func (in *RegistryList) DeepCopy() *RegistryList {
	if in == nil {
		return nil
	}
	out := new(RegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryLogin) DeepCopyInto(out *RegistryLogin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryLogin.
func (in *RegistryLogin) DeepCopy() *RegistryLogin {
	if in == nil {
		return nil
	}
	out := new(RegistryLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryPVC) DeepCopyInto(out *RegistryPVC) {
	*out = *in
	if in.Exist != nil {
		in, out := &in.Exist, &out.Exist
		*out = new(ExistPvc)
		**out = **in
	}
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(CreatePvc)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryPVC.
func (in *RegistryPVC) DeepCopy() *RegistryPVC {
	if in == nil {
		return nil
	}
	out := new(RegistryPVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryService) DeepCopyInto(out *RegistryService) {
	*out = *in
	out.Ingress = in.Ingress
	out.LoadBalancer = in.LoadBalancer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryService.
func (in *RegistryService) DeepCopy() *RegistryService {
	if in == nil {
		return nil
	}
	out := new(RegistryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
	in.ReplicaSet.DeepCopyInto(&out.ReplicaSet)
	in.RegistryDeployment.DeepCopyInto(&out.RegistryDeployment)
	out.RegistryService = in.RegistryService
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStatus) DeepCopyInto(out *RegistryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PhaseChangedAt.DeepCopyInto(&out.PhaseChangedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStatus.
func (in *RegistryStatus) DeepCopy() *RegistryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerKey) DeepCopyInto(out *SignerKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerKey.
func (in *SignerKey) DeepCopy() *SignerKey {
	if in == nil {
		return nil
	}
	out := new(SignerKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignerKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerKeyList) DeepCopyInto(out *SignerKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SignerKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerKeyList.
func (in *SignerKeyList) DeepCopy() *SignerKeyList {
	if in == nil {
		return nil
	}
	out := new(SignerKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignerKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerKeySpec) DeepCopyInto(out *SignerKeySpec) {
	*out = *in
	out.Root = in.Root
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make(map[string]TrustKey, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerKeySpec.
func (in *SignerKeySpec) DeepCopy() *SignerKeySpec {
	if in == nil {
		return nil
	}
	out := new(SignerKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerKeyStatus) DeepCopyInto(out *SignerKeyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerKeyStatus.
func (in *SignerKeyStatus) DeepCopy() *SignerKeyStatus {
	if in == nil {
		return nil
	}
	out := new(SignerKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustKey) DeepCopyInto(out *TrustKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustKey.
func (in *TrustKey) DeepCopy() *TrustKey {
	if in == nil {
		return nil
	}
	out := new(TrustKey)
	in.DeepCopyInto(out)
	return out
}
