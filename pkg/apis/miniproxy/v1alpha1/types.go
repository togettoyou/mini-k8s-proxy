package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProxyRoute is a specification for a ProxyRoute resource
type ProxyRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProxyRouteSpec `json:"spec"`
}

// ProxyRouteSpec is the spec for a ProxyRoute resource
type ProxyRouteSpec struct {
	ServiceName string `json:"serviceName"`
	Namespace   string `json:"namespace,omitempty"`
	Port        int32  `json:"port,omitempty"`
	Scheme      string `json:"scheme,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProxyRouteList is a list of ProxyRoute resources
type ProxyRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ProxyRoute `json:"items"`
}
