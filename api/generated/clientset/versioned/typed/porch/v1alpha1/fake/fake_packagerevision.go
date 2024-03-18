// Copyright 2023 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/nephio-project/porch/api/porch/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePackageRevisions implements PackageRevisionInterface
type FakePackageRevisions struct {
	Fake *FakePorchV1alpha1
	ns   string
}

var packagerevisionsResource = v1alpha1.SchemeGroupVersion.WithResource("packagerevisions")

var packagerevisionsKind = v1alpha1.SchemeGroupVersion.WithKind("PackageRevision")

// Get takes name of the packageRevision, and returns the corresponding packageRevision object, and an error if there is any.
func (c *FakePackageRevisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PackageRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(packagerevisionsResource, c.ns, name), &v1alpha1.PackageRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevision), err
}

// List takes label and field selectors, and returns the list of PackageRevisions that match those selectors.
func (c *FakePackageRevisions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PackageRevisionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(packagerevisionsResource, packagerevisionsKind, c.ns, opts), &v1alpha1.PackageRevisionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PackageRevisionList{ListMeta: obj.(*v1alpha1.PackageRevisionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PackageRevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested packageRevisions.
func (c *FakePackageRevisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(packagerevisionsResource, c.ns, opts))

}

// Create takes the representation of a packageRevision and creates it.  Returns the server's representation of the packageRevision, and an error, if there is any.
func (c *FakePackageRevisions) Create(ctx context.Context, packageRevision *v1alpha1.PackageRevision, opts v1.CreateOptions) (result *v1alpha1.PackageRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(packagerevisionsResource, c.ns, packageRevision), &v1alpha1.PackageRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevision), err
}

// Update takes the representation of a packageRevision and updates it. Returns the server's representation of the packageRevision, and an error, if there is any.
func (c *FakePackageRevisions) Update(ctx context.Context, packageRevision *v1alpha1.PackageRevision, opts v1.UpdateOptions) (result *v1alpha1.PackageRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(packagerevisionsResource, c.ns, packageRevision), &v1alpha1.PackageRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevision), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePackageRevisions) UpdateStatus(ctx context.Context, packageRevision *v1alpha1.PackageRevision, opts v1.UpdateOptions) (*v1alpha1.PackageRevision, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(packagerevisionsResource, "status", c.ns, packageRevision), &v1alpha1.PackageRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevision), err
}

// Delete takes name of the packageRevision and deletes it. Returns an error if one occurs.
func (c *FakePackageRevisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(packagerevisionsResource, c.ns, name, opts), &v1alpha1.PackageRevision{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePackageRevisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(packagerevisionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PackageRevisionList{})
	return err
}

// Patch applies the patch and returns the patched packageRevision.
func (c *FakePackageRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PackageRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(packagerevisionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PackageRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevision), err
}

// UpdateApproval takes the representation of a packageRevision and updates it. Returns the server's representation of the packageRevision, and an error, if there is any.
func (c *FakePackageRevisions) UpdateApproval(ctx context.Context, packageRevisionName string, packageRevision *v1alpha1.PackageRevision, opts v1.UpdateOptions) (result *v1alpha1.PackageRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(packagerevisionsResource, "approval", c.ns, packageRevision), &v1alpha1.PackageRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRevision), err
}
