//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/nephio-project/porch/pkg/kpt/api/kptfile/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Downstream) DeepCopyInto(out *Downstream) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Downstream.
func (in *Downstream) DeepCopy() *Downstream {
	if in == nil {
		return nil
	}
	out := new(Downstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownstreamTarget) DeepCopyInto(out *DownstreamTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownstreamTarget.
func (in *DownstreamTarget) DeepCopy() *DownstreamTarget {
	if in == nil {
		return nil
	}
	out := new(DownstreamTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InjectionSelector) DeepCopyInto(out *InjectionSelector) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InjectionSelector.
func (in *InjectionSelector) DeepCopy() *InjectionSelector {
	if in == nil {
		return nil
	}
	out := new(InjectionSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageContext) DeepCopyInto(out *PackageContext) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RemoveKeys != nil {
		in, out := &in.RemoveKeys, &out.RemoveKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageContext.
func (in *PackageContext) DeepCopy() *PackageContext {
	if in == nil {
		return nil
	}
	out := new(PackageContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageVariant) DeepCopyInto(out *PackageVariant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageVariant.
func (in *PackageVariant) DeepCopy() *PackageVariant {
	if in == nil {
		return nil
	}
	out := new(PackageVariant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageVariant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageVariantList) DeepCopyInto(out *PackageVariantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageVariant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageVariantList.
func (in *PackageVariantList) DeepCopy() *PackageVariantList {
	if in == nil {
		return nil
	}
	out := new(PackageVariantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageVariantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageVariantSpec) DeepCopyInto(out *PackageVariantSpec) {
	*out = *in
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = new(Upstream)
		**out = **in
	}
	if in.Downstream != nil {
		in, out := &in.Downstream, &out.Downstream
		*out = new(Downstream)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PackageContext != nil {
		in, out := &in.PackageContext, &out.PackageContext
		*out = new(PackageContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = new(v1.Pipeline)
		(*in).DeepCopyInto(*out)
	}
	if in.Injectors != nil {
		in, out := &in.Injectors, &out.Injectors
		*out = make([]InjectionSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageVariantSpec.
func (in *PackageVariantSpec) DeepCopy() *PackageVariantSpec {
	if in == nil {
		return nil
	}
	out := new(PackageVariantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageVariantStatus) DeepCopyInto(out *PackageVariantStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DownstreamTargets != nil {
		in, out := &in.DownstreamTargets, &out.DownstreamTargets
		*out = make([]DownstreamTarget, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageVariantStatus.
func (in *PackageVariantStatus) DeepCopy() *PackageVariantStatus {
	if in == nil {
		return nil
	}
	out := new(PackageVariantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Upstream) DeepCopyInto(out *Upstream) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Upstream.
func (in *Upstream) DeepCopy() *Upstream {
	if in == nil {
		return nil
	}
	out := new(Upstream)
	in.DeepCopyInto(out)
	return out
}
