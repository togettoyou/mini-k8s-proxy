package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MiniProxy is a specification for a MiniProxy resource
type MiniProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MiniProxySpec `json:"spec"`
}

// MiniProxySpec is the spec for a MiniProxy resource
type MiniProxySpec struct {
	ServiceName string             `json:"serviceName"`
	Namespace   string             `json:"namespace,omitempty"`
	Port        intstr.IntOrString `json:"port,omitempty"`
	Scheme      string             `json:"scheme,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MiniProxyList is a list of MiniProxy resources
type MiniProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MiniProxy `json:"items"`
}
