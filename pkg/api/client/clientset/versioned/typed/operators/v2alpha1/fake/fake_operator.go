/*
Copyright Red Hat, Inc.

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

package fake

import (
	"context"

	v2alpha1 "github.com/operator-framework/api/pkg/operators/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperators implements OperatorInterface
type FakeOperators struct {
	Fake *FakeOperatorsV2alpha1
}

var operatorsResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "v2alpha1", Resource: "operators"}

var operatorsKind = schema.GroupVersionKind{Group: "operators.coreos.com", Version: "v2alpha1", Kind: "Operator"}

// Get takes name of the operator, and returns the corresponding operator object, and an error if there is any.
func (c *FakeOperators) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(operatorsResource, name), &v2alpha1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.Operator), err
}

// List takes label and field selectors, and returns the list of Operators that match those selectors.
func (c *FakeOperators) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.OperatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(operatorsResource, operatorsKind, opts), &v2alpha1.OperatorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.OperatorList{ListMeta: obj.(*v2alpha1.OperatorList).ListMeta}
	for _, item := range obj.(*v2alpha1.OperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operators.
func (c *FakeOperators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(operatorsResource, opts))
}

// Create takes the representation of a operator and creates it.  Returns the server's representation of the operator, and an error, if there is any.
func (c *FakeOperators) Create(ctx context.Context, operator *v2alpha1.Operator, opts v1.CreateOptions) (result *v2alpha1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(operatorsResource, operator), &v2alpha1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.Operator), err
}

// Update takes the representation of a operator and updates it. Returns the server's representation of the operator, and an error, if there is any.
func (c *FakeOperators) Update(ctx context.Context, operator *v2alpha1.Operator, opts v1.UpdateOptions) (result *v2alpha1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(operatorsResource, operator), &v2alpha1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.Operator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperators) UpdateStatus(ctx context.Context, operator *v2alpha1.Operator, opts v1.UpdateOptions) (*v2alpha1.Operator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(operatorsResource, "status", operator), &v2alpha1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.Operator), err
}

// Delete takes name of the operator and deletes it. Returns an error if one occurs.
func (c *FakeOperators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(operatorsResource, name), &v2alpha1.Operator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(operatorsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.OperatorList{})
	return err
}

// Patch applies the patch and returns the patched operator.
func (c *FakeOperators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(operatorsResource, name, pt, data, subresources...), &v2alpha1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.Operator), err
}