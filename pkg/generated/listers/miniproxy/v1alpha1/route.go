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

// RouteLister helps list Routes.
// All objects returned here must be treated as read-only.
type RouteLister interface {
	// List lists all Routes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Route, err error)
	// Routes returns an object that can list and get Routes.
	Routes(namespace string) RouteNamespaceLister
	RouteListerExpansion
}

// routeLister implements the RouteLister interface.
type routeLister struct {
	indexer cache.Indexer
}

// NewRouteLister returns a new RouteLister.
func NewRouteLister(indexer cache.Indexer) RouteLister {
	return &routeLister{indexer: indexer}
}

// List lists all Routes in the indexer.
func (s *routeLister) List(selector labels.Selector) (ret []*v1alpha1.Route, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route))
	})
	return ret, err
}

// Routes returns an object that can list and get Routes.
func (s *routeLister) Routes(namespace string) RouteNamespaceLister {
	return routeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RouteNamespaceLister helps list and get Routes.
// All objects returned here must be treated as read-only.
type RouteNamespaceLister interface {
	// List lists all Routes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Route, err error)
	// Get retrieves the Route from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Route, error)
	RouteNamespaceListerExpansion
}

// routeNamespaceLister implements the RouteNamespaceLister
// interface.
type routeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Routes in the indexer for a given namespace.
func (s routeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Route, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route))
	})
	return ret, err
}

// Get retrieves the Route from the indexer for a given namespace and name.
func (s routeNamespaceLister) Get(name string) (*v1alpha1.Route, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("route"), name)
	}
	return obj.(*v1alpha1.Route), nil
}