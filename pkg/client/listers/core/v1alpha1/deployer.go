/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package v1alpha1

import (
	v1alpha1 "github.com/projectriff/system/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeployerLister helps list Deployers.
type DeployerLister interface {
	// List lists all Deployers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Deployer, err error)
	// Deployers returns an object that can list and get Deployers.
	Deployers(namespace string) DeployerNamespaceLister
	DeployerListerExpansion
}

// deployerLister implements the DeployerLister interface.
type deployerLister struct {
	indexer cache.Indexer
}

// NewDeployerLister returns a new DeployerLister.
func NewDeployerLister(indexer cache.Indexer) DeployerLister {
	return &deployerLister{indexer: indexer}
}

// List lists all Deployers in the indexer.
func (s *deployerLister) List(selector labels.Selector) (ret []*v1alpha1.Deployer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Deployer))
	})
	return ret, err
}

// Deployers returns an object that can list and get Deployers.
func (s *deployerLister) Deployers(namespace string) DeployerNamespaceLister {
	return deployerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeployerNamespaceLister helps list and get Deployers.
type DeployerNamespaceLister interface {
	// List lists all Deployers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Deployer, err error)
	// Get retrieves the Deployer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Deployer, error)
	DeployerNamespaceListerExpansion
}

// deployerNamespaceLister implements the DeployerNamespaceLister
// interface.
type deployerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Deployers in the indexer for a given namespace.
func (s deployerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Deployer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Deployer))
	})
	return ret, err
}

// Get retrieves the Deployer from the indexer for a given namespace and name.
func (s deployerNamespaceLister) Get(name string) (*v1alpha1.Deployer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("deployer"), name)
	}
	return obj.(*v1alpha1.Deployer), nil
}
