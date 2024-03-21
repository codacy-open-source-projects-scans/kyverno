/*
Copyright The Kubernetes Authors.

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

	v2 "github.com/kyverno/kyverno/api/kyverno/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCleanupPolicies implements CleanupPolicyInterface
type FakeCleanupPolicies struct {
	Fake *FakeKyvernoV2
	ns   string
}

var cleanuppoliciesResource = v2.SchemeGroupVersion.WithResource("cleanuppolicies")

var cleanuppoliciesKind = v2.SchemeGroupVersion.WithKind("CleanupPolicy")

// Get takes name of the cleanupPolicy, and returns the corresponding cleanupPolicy object, and an error if there is any.
func (c *FakeCleanupPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CleanupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cleanuppoliciesResource, c.ns, name), &v2.CleanupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CleanupPolicy), err
}

// List takes label and field selectors, and returns the list of CleanupPolicies that match those selectors.
func (c *FakeCleanupPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2.CleanupPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cleanuppoliciesResource, cleanuppoliciesKind, c.ns, opts), &v2.CleanupPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.CleanupPolicyList{ListMeta: obj.(*v2.CleanupPolicyList).ListMeta}
	for _, item := range obj.(*v2.CleanupPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cleanupPolicies.
func (c *FakeCleanupPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cleanuppoliciesResource, c.ns, opts))

}

// Create takes the representation of a cleanupPolicy and creates it.  Returns the server's representation of the cleanupPolicy, and an error, if there is any.
func (c *FakeCleanupPolicies) Create(ctx context.Context, cleanupPolicy *v2.CleanupPolicy, opts v1.CreateOptions) (result *v2.CleanupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cleanuppoliciesResource, c.ns, cleanupPolicy), &v2.CleanupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CleanupPolicy), err
}

// Update takes the representation of a cleanupPolicy and updates it. Returns the server's representation of the cleanupPolicy, and an error, if there is any.
func (c *FakeCleanupPolicies) Update(ctx context.Context, cleanupPolicy *v2.CleanupPolicy, opts v1.UpdateOptions) (result *v2.CleanupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cleanuppoliciesResource, c.ns, cleanupPolicy), &v2.CleanupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CleanupPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCleanupPolicies) UpdateStatus(ctx context.Context, cleanupPolicy *v2.CleanupPolicy, opts v1.UpdateOptions) (*v2.CleanupPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cleanuppoliciesResource, "status", c.ns, cleanupPolicy), &v2.CleanupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CleanupPolicy), err
}

// Delete takes name of the cleanupPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCleanupPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cleanuppoliciesResource, c.ns, name, opts), &v2.CleanupPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCleanupPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cleanuppoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.CleanupPolicyList{})
	return err
}

// Patch applies the patch and returns the patched cleanupPolicy.
func (c *FakeCleanupPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CleanupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cleanuppoliciesResource, c.ns, name, pt, data, subresources...), &v2.CleanupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.CleanupPolicy), err
}
