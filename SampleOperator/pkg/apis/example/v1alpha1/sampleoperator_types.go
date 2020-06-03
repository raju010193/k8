package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SampleOperatorSpec defines the desired state of SampleOperator
type SampleOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
    Replicas *int32 `json:"replicas"`
	Host      string `json:"host"`
	Recipient string `json:"recipient"`
}

// SampleOperatorStatus defines the observed state of SampleOperator
type SampleOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SampleOperator is the Schema for the sampleoperators API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=sampleoperators,scope=Namespaced
type SampleOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SampleOperatorSpec   `json:"spec,omitempty"`
	Status SampleOperatorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SampleOperatorList contains a list of SampleOperator
type SampleOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SampleOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SampleOperator{}, &SampleOperatorList{},&SampleOperatorSpec{})
}
