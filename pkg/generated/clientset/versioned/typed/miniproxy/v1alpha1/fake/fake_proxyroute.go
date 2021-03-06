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

package fake

import (
	"context"
	v1alpha1 "mini-k8s-proxy/pkg/apis/miniproxy/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProxyRoutes implements ProxyRouteInterface
type FakeProxyRoutes struct {
	Fake *FakeMiniproxyV1alpha1
	ns   string
}

var proxyroutesResource = schema.GroupVersionResource{Group: "miniproxy.togettoyou.com", Version: "v1alpha1", Resource: "proxyroutes"}

var proxyroutesKind = schema.GroupVersionKind{Group: "miniproxy.togettoyou.com", Version: "v1alpha1", Kind: "ProxyRoute"}

// Get takes name of the proxyRoute, and returns the corresponding proxyRoute object, and an error if there is any.
func (c *FakeProxyRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProxyRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(proxyroutesResource, c.ns, name), &v1alpha1.ProxyRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyRoute), err
}

// List takes label and field selectors, and returns the list of ProxyRoutes that match those selectors.
func (c *FakeProxyRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProxyRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(proxyroutesResource, proxyroutesKind, c.ns, opts), &v1alpha1.ProxyRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProxyRouteList{ListMeta: obj.(*v1alpha1.ProxyRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProxyRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested proxyRoutes.
func (c *FakeProxyRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(proxyroutesResource, c.ns, opts))

}

// Create takes the representation of a proxyRoute and creates it.  Returns the server's representation of the proxyRoute, and an error, if there is any.
func (c *FakeProxyRoutes) Create(ctx context.Context, proxyRoute *v1alpha1.ProxyRoute, opts v1.CreateOptions) (result *v1alpha1.ProxyRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(proxyroutesResource, c.ns, proxyRoute), &v1alpha1.ProxyRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyRoute), err
}

// Update takes the representation of a proxyRoute and updates it. Returns the server's representation of the proxyRoute, and an error, if there is any.
func (c *FakeProxyRoutes) Update(ctx context.Context, proxyRoute *v1alpha1.ProxyRoute, opts v1.UpdateOptions) (result *v1alpha1.ProxyRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(proxyroutesResource, c.ns, proxyRoute), &v1alpha1.ProxyRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyRoute), err
}

// Delete takes name of the proxyRoute and deletes it. Returns an error if one occurs.
func (c *FakeProxyRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(proxyroutesResource, c.ns, name), &v1alpha1.ProxyRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProxyRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(proxyroutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProxyRouteList{})
	return err
}

// Patch applies the patch and returns the patched proxyRoute.
func (c *FakeProxyRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProxyRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(proxyroutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProxyRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyRoute), err
}
