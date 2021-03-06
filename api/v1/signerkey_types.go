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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SignerKeySpec defines the desired state of SignerKey
type SignerKeySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SignerKey. Edit SignerKey_types.go to remove/update
	Root TrustKey `json:"root,omitempty"`
	// Targets is {namespace/registryName/imageName: TrustKey{}, ...}
	Targets map[string]TrustKey `json:"targets,omitempty"`
}

// TrustKey defines key and value set
type TrustKey struct {
	ID         string `json:"id"`
	Key        string `json:"key"`
	PassPhrase string `json:"passPhrase"`
}

// SignerKeyStatus defines the observed state of SignerKey
type SignerKeyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=sk

// SignerKey is the Schema for the signerkeys API
type SignerKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SignerKeySpec   `json:"spec,omitempty"`
	Status SignerKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SignerKeyList contains a list of SignerKey
type SignerKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SignerKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SignerKey{}, &SignerKeyList{})
}
