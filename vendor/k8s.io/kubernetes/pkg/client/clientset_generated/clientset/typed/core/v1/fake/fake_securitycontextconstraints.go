/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1 "k8s.io/kubernetes/pkg/api/v1"
)

// FakeSecurityContextConstraints implements SecurityContextConstraintsInterface
type FakeSecurityContextConstraints struct {
	Fake *FakeCoreV1
}

var securitycontextconstraintsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "securitycontextconstraints"}

func (c *FakeSecurityContextConstraints) Create(securityContextConstraints *v1.SecurityContextConstraints) (result *v1.SecurityContextConstraints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(securitycontextconstraintsResource, securityContextConstraints), &v1.SecurityContextConstraints{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SecurityContextConstraints), err
}

func (c *FakeSecurityContextConstraints) Update(securityContextConstraints *v1.SecurityContextConstraints) (result *v1.SecurityContextConstraints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(securitycontextconstraintsResource, securityContextConstraints), &v1.SecurityContextConstraints{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SecurityContextConstraints), err
}

func (c *FakeSecurityContextConstraints) Delete(name string, options *meta_v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(securitycontextconstraintsResource, name), &v1.SecurityContextConstraints{})
	return err
}

func (c *FakeSecurityContextConstraints) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(securitycontextconstraintsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1.SecurityContextConstraintsList{})
	return err
}

func (c *FakeSecurityContextConstraints) Get(name string, options meta_v1.GetOptions) (result *v1.SecurityContextConstraints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(securitycontextconstraintsResource, name), &v1.SecurityContextConstraints{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SecurityContextConstraints), err
}

func (c *FakeSecurityContextConstraints) List(opts meta_v1.ListOptions) (result *v1.SecurityContextConstraintsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(securitycontextconstraintsResource, opts), &v1.SecurityContextConstraintsList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SecurityContextConstraintsList{}
	for _, item := range obj.(*v1.SecurityContextConstraintsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested securityContextConstraints.
func (c *FakeSecurityContextConstraints) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(securitycontextconstraintsResource, opts))
}

// Patch applies the patch and returns the patched securityContextConstraints.
func (c *FakeSecurityContextConstraints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SecurityContextConstraints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(securitycontextconstraintsResource, name, data, subresources...), &v1.SecurityContextConstraints{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SecurityContextConstraints), err
}
