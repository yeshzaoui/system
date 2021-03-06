/*
Copyright 2019 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/projectriff/system/pkg/apis/streaming/v1alpha1"
	scheme "github.com/projectriff/system/pkg/client/clientset/versioned/scheme"
)

// KafkaGatewaysGetter has a method to return a KafkaGatewayInterface.
// A group's client should implement this interface.
type KafkaGatewaysGetter interface {
	KafkaGateways(namespace string) KafkaGatewayInterface
}

// KafkaGatewayInterface has methods to work with KafkaGateway resources.
type KafkaGatewayInterface interface {
	Create(*v1alpha1.KafkaGateway) (*v1alpha1.KafkaGateway, error)
	Update(*v1alpha1.KafkaGateway) (*v1alpha1.KafkaGateway, error)
	UpdateStatus(*v1alpha1.KafkaGateway) (*v1alpha1.KafkaGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KafkaGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.KafkaGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaGateway, err error)
	KafkaGatewayExpansion
}

// kafkaGateways implements KafkaGatewayInterface
type kafkaGateways struct {
	client rest.Interface
	ns     string
}

// newKafkaGateways returns a KafkaGateways
func newKafkaGateways(c *StreamingV1alpha1Client, namespace string) *kafkaGateways {
	return &kafkaGateways{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaGateway, and returns the corresponding kafkaGateway object, and an error if there is any.
func (c *kafkaGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.KafkaGateway, err error) {
	result = &v1alpha1.KafkaGateway{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkagateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaGateways that match those selectors.
func (c *kafkaGateways) List(opts v1.ListOptions) (result *v1alpha1.KafkaGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KafkaGatewayList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkagateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaGateways.
func (c *kafkaGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkagateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kafkaGateway and creates it.  Returns the server's representation of the kafkaGateway, and an error, if there is any.
func (c *kafkaGateways) Create(kafkaGateway *v1alpha1.KafkaGateway) (result *v1alpha1.KafkaGateway, err error) {
	result = &v1alpha1.KafkaGateway{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kafkagateways").
		Body(kafkaGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kafkaGateway and updates it. Returns the server's representation of the kafkaGateway, and an error, if there is any.
func (c *kafkaGateways) Update(kafkaGateway *v1alpha1.KafkaGateway) (result *v1alpha1.KafkaGateway, err error) {
	result = &v1alpha1.KafkaGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkagateways").
		Name(kafkaGateway.Name).
		Body(kafkaGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kafkaGateways) UpdateStatus(kafkaGateway *v1alpha1.KafkaGateway) (result *v1alpha1.KafkaGateway, err error) {
	result = &v1alpha1.KafkaGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkagateways").
		Name(kafkaGateway.Name).
		SubResource("status").
		Body(kafkaGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the kafkaGateway and deletes it. Returns an error if one occurs.
func (c *kafkaGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkagateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kafkaGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkagateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kafkaGateway.
func (c *kafkaGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KafkaGateway, err error) {
	result = &v1alpha1.KafkaGateway{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kafkagateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
