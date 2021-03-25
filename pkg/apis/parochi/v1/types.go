package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Greeting is a top-level type
type Greeting struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status GreetingStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec GreetingSpec `json:"spec,omitempty"`
}

// custom spec
type GreetingSpec struct {
	Name string `json:"name,omitempty"`
}

// custom status
type GreetingStatus struct {
	status string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type GreetingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Greeting `json:"items"`
}
