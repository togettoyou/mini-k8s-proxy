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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "mini-k8s-proxy/pkg/apis/miniproxy/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MiniProxyLister helps list MiniProxies.
// All objects returned here must be treated as read-only.
type MiniProxyLister interface {
	// List lists all MiniProxies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MiniProxy, err error)
	// MiniProxies returns an object that can list and get MiniProxies.
	MiniProxies(namespace string) MiniProxyNamespaceLister
	MiniProxyListerExpansion
}

// miniProxyLister implements the MiniProxyLister interface.
type miniProxyLister struct {
	indexer cache.Indexer
}

// NewMiniProxyLister returns a new MiniProxyLister.
func NewMiniProxyLister(indexer cache.Indexer) MiniProxyLister {
	return &miniProxyLister{indexer: indexer}
}

// List lists all MiniProxies in the indexer.
func (s *miniProxyLister) List(selector labels.Selector) (ret []*v1alpha1.MiniProxy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MiniProxy))
	})
	return ret, err
}

// MiniProxies returns an object that can list and get MiniProxies.
func (s *miniProxyLister) MiniProxies(namespace string) MiniProxyNamespaceLister {
	return miniProxyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MiniProxyNamespaceLister helps list and get MiniProxies.
// All objects returned here must be treated as read-only.
type MiniProxyNamespaceLister interface {
	// List lists all MiniProxies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MiniProxy, err error)
	// Get retrieves the MiniProxy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MiniProxy, error)
	MiniProxyNamespaceListerExpansion
}

// miniProxyNamespaceLister implements the MiniProxyNamespaceLister
// interface.
type miniProxyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MiniProxies in the indexer for a given namespace.
func (s miniProxyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MiniProxy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MiniProxy))
	})
	return ret, err
}

// Get retrieves the MiniProxy from the indexer for a given namespace and name.
func (s miniProxyNamespaceLister) Get(name string) (*v1alpha1.MiniProxy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("miniproxy"), name)
	}
	return obj.(*v1alpha1.MiniProxy), nil
}