/*
Copyright 2023.

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

// EngineerSpec defines the desired state of Engineer
type EngineerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Job     string `json:"job,omitempty"`
	Company string `json:"company,omitempty"`
}

// EngineerStatus defines the observed state of Engineer
type EngineerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Happy bool `json:"happy"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Engineer is the Schema for the engineers API
type Engineer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EngineerSpec   `json:"spec,omitempty"`
	Status EngineerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EngineerList contains a list of Engineer
type EngineerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Engineer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Engineer{}, &EngineerList{})
}
