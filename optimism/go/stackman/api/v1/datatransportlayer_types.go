/*
Copyright 2021.

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DataTransportLayerSpec defines the desired state of DataTransportLayer
type DataTransportLayerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image                  string      `json:"image,omitempty"`
	L1URL                  string      `json:"l1_url,omitempty"`
	L1TimeoutSeconds       int         `json:"l1_timeout_seconds,omitempty"`
	DeployerURL            string      `json:"deployer_url,omitempty"`
	DeployerTimeoutSeconds int         `json:"deployer_timeout_seconds,omitempty"`
	DataPVC                *PVCConfig  `json:"data_pvc,omitempty"`
	Env                    []v1.EnvVar `json:"env,omitempty"`
}

// DataTransportLayerStatus defines the observed state of DataTransportLayer
type DataTransportLayerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DataTransportLayer is the Schema for the datatransportlayers API
type DataTransportLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataTransportLayerSpec   `json:"spec,omitempty"`
	Status DataTransportLayerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DataTransportLayerList contains a list of DataTransportLayer
type DataTransportLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataTransportLayer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataTransportLayer{}, &DataTransportLayerList{})
}
