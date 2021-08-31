/*
MIT License

Copyright (c) 2021 togettoyou

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "mini-k8s-proxy/pkg/apis/miniproxy/v1alpha1"
	"mini-k8s-proxy/pkg/generated/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type MiniproxyV1alpha1Interface interface {
	RESTClient() rest.Interface
	RoutesGetter
}

// MiniproxyV1alpha1Client is used to interact with features provided by the miniproxy.togettoyou.com group.
type MiniproxyV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MiniproxyV1alpha1Client) Routes(namespace string) RouteInterface {
	return newRoutes(c, namespace)
}

// NewForConfig creates a new MiniproxyV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MiniproxyV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MiniproxyV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MiniproxyV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MiniproxyV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MiniproxyV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MiniproxyV1alpha1Client {
	return &MiniproxyV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MiniproxyV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
