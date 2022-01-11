package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AlertOverridesSpec defines the desired state of AlertOverrides
type AlertOverridesSpec struct {
	// +required
	Name string `json:"name"`
}

// AlertOverridesStatus defines the observed state of AlertOverrides.
// It should always be reconstructable from the state of the cluster and/or outside world.
type AlertOverridesStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of cluster
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertOverrides is the Schema for the alertoverrides API
// +k8s:openapi-gen=true
type AlertOverrides struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertOverridesSpec   `json:"spec,omitempty"`
	Status AlertOverridesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertOverridesList contains a list of AlertOverrides
type AlertOverridesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertOverrides `json:"items"`
}
