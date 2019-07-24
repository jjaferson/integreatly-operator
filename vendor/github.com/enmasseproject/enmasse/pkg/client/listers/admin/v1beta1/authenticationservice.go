/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/enmasseproject/enmasse/pkg/apis/admin/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuthenticationServiceLister helps list AuthenticationServices.
type AuthenticationServiceLister interface {
	// List lists all AuthenticationServices in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.AuthenticationService, err error)
	// AuthenticationServices returns an object that can list and get AuthenticationServices.
	AuthenticationServices(namespace string) AuthenticationServiceNamespaceLister
	AuthenticationServiceListerExpansion
}

// authenticationServiceLister implements the AuthenticationServiceLister interface.
type authenticationServiceLister struct {
	indexer cache.Indexer
}

// NewAuthenticationServiceLister returns a new AuthenticationServiceLister.
func NewAuthenticationServiceLister(indexer cache.Indexer) AuthenticationServiceLister {
	return &authenticationServiceLister{indexer: indexer}
}

// List lists all AuthenticationServices in the indexer.
func (s *authenticationServiceLister) List(selector labels.Selector) (ret []*v1beta1.AuthenticationService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.AuthenticationService))
	})
	return ret, err
}

// AuthenticationServices returns an object that can list and get AuthenticationServices.
func (s *authenticationServiceLister) AuthenticationServices(namespace string) AuthenticationServiceNamespaceLister {
	return authenticationServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuthenticationServiceNamespaceLister helps list and get AuthenticationServices.
type AuthenticationServiceNamespaceLister interface {
	// List lists all AuthenticationServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.AuthenticationService, err error)
	// Get retrieves the AuthenticationService from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.AuthenticationService, error)
	AuthenticationServiceNamespaceListerExpansion
}

// authenticationServiceNamespaceLister implements the AuthenticationServiceNamespaceLister
// interface.
type authenticationServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuthenticationServices in the indexer for a given namespace.
func (s authenticationServiceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.AuthenticationService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.AuthenticationService))
	})
	return ret, err
}

// Get retrieves the AuthenticationService from the indexer for a given namespace and name.
func (s authenticationServiceNamespaceLister) Get(name string) (*v1beta1.AuthenticationService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("authenticationservice"), name)
	}
	return obj.(*v1beta1.AuthenticationService), nil
}