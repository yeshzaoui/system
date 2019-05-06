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
	v1alpha1 "github.com/projectriff/system/pkg/apis/stream/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StreamLister helps list Streams.
type StreamLister interface {
	// List lists all Streams in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Stream, err error)
	// Streams returns an object that can list and get Streams.
	Streams(namespace string) StreamNamespaceLister
	StreamListerExpansion
}

// streamLister implements the StreamLister interface.
type streamLister struct {
	indexer cache.Indexer
}

// NewStreamLister returns a new StreamLister.
func NewStreamLister(indexer cache.Indexer) StreamLister {
	return &streamLister{indexer: indexer}
}

// List lists all Streams in the indexer.
func (s *streamLister) List(selector labels.Selector) (ret []*v1alpha1.Stream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Stream))
	})
	return ret, err
}

// Streams returns an object that can list and get Streams.
func (s *streamLister) Streams(namespace string) StreamNamespaceLister {
	return streamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamNamespaceLister helps list and get Streams.
type StreamNamespaceLister interface {
	// List lists all Streams in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Stream, err error)
	// Get retrieves the Stream from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Stream, error)
	StreamNamespaceListerExpansion
}

// streamNamespaceLister implements the StreamNamespaceLister
// interface.
type streamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Streams in the indexer for a given namespace.
func (s streamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Stream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Stream))
	})
	return ret, err
}

// Get retrieves the Stream from the indexer for a given namespace and name.
func (s streamNamespaceLister) Get(name string) (*v1alpha1.Stream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("stream"), name)
	}
	return obj.(*v1alpha1.Stream), nil
}
