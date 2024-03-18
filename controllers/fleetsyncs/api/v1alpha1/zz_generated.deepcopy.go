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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembership) DeepCopyInto(out *FleetMembership) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Data.DeepCopyInto(&out.Data)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembership.
func (in *FleetMembership) DeepCopy() *FleetMembership {
	if in == nil {
		return nil
	}
	out := new(FleetMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetMembership) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembershipBinding) DeepCopyInto(out *FleetMembershipBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Data.DeepCopyInto(&out.Data)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembershipBinding.
func (in *FleetMembershipBinding) DeepCopy() *FleetMembershipBinding {
	if in == nil {
		return nil
	}
	out := new(FleetMembershipBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetMembershipBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembershipBindingData) DeepCopyInto(out *FleetMembershipBindingData) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.State = in.State
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembershipBindingData.
func (in *FleetMembershipBindingData) DeepCopy() *FleetMembershipBindingData {
	if in == nil {
		return nil
	}
	out := new(FleetMembershipBindingData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembershipBindingList) DeepCopyInto(out *FleetMembershipBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetMembershipBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembershipBindingList.
func (in *FleetMembershipBindingList) DeepCopy() *FleetMembershipBindingList {
	if in == nil {
		return nil
	}
	out := new(FleetMembershipBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetMembershipBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembershipBindingStatus) DeepCopyInto(out *FleetMembershipBindingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembershipBindingStatus.
func (in *FleetMembershipBindingStatus) DeepCopy() *FleetMembershipBindingStatus {
	if in == nil {
		return nil
	}
	out := new(FleetMembershipBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembershipData) DeepCopyInto(out *FleetMembershipData) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.State = in.State
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembershipData.
func (in *FleetMembershipData) DeepCopy() *FleetMembershipData {
	if in == nil {
		return nil
	}
	out := new(FleetMembershipData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembershipList) DeepCopyInto(out *FleetMembershipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembershipList.
func (in *FleetMembershipList) DeepCopy() *FleetMembershipList {
	if in == nil {
		return nil
	}
	out := new(FleetMembershipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetMembershipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetMembershipStatus) DeepCopyInto(out *FleetMembershipStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetMembershipStatus.
func (in *FleetMembershipStatus) DeepCopy() *FleetMembershipStatus {
	if in == nil {
		return nil
	}
	out := new(FleetMembershipStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetScope) DeepCopyInto(out *FleetScope) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Data.DeepCopyInto(&out.Data)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetScope.
func (in *FleetScope) DeepCopy() *FleetScope {
	if in == nil {
		return nil
	}
	out := new(FleetScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetScope) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetScopeData) DeepCopyInto(out *FleetScopeData) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.State = in.State
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetScopeData.
func (in *FleetScopeData) DeepCopy() *FleetScopeData {
	if in == nil {
		return nil
	}
	out := new(FleetScopeData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetScopeList) DeepCopyInto(out *FleetScopeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetScope, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetScopeList.
func (in *FleetScopeList) DeepCopy() *FleetScopeList {
	if in == nil {
		return nil
	}
	out := new(FleetScopeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetScopeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetScopeStatus) DeepCopyInto(out *FleetScopeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetScopeStatus.
func (in *FleetScopeStatus) DeepCopy() *FleetScopeStatus {
	if in == nil {
		return nil
	}
	out := new(FleetScopeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSync) DeepCopyInto(out *FleetSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSync.
func (in *FleetSync) DeepCopy() *FleetSync {
	if in == nil {
		return nil
	}
	out := new(FleetSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSyncList) DeepCopyInto(out *FleetSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSyncList.
func (in *FleetSyncList) DeepCopy() *FleetSyncList {
	if in == nil {
		return nil
	}
	out := new(FleetSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSyncSpec) DeepCopyInto(out *FleetSyncSpec) {
	*out = *in
	if in.ProjectIds != nil {
		in, out := &in.ProjectIds, &out.ProjectIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSyncSpec.
func (in *FleetSyncSpec) DeepCopy() *FleetSyncSpec {
	if in == nil {
		return nil
	}
	out := new(FleetSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSyncStatus) DeepCopyInto(out *FleetSyncStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSyncStatus.
func (in *FleetSyncStatus) DeepCopy() *FleetSyncStatus {
	if in == nil {
		return nil
	}
	out := new(FleetSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipBindingState) DeepCopyInto(out *MembershipBindingState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipBindingState.
func (in *MembershipBindingState) DeepCopy() *MembershipBindingState {
	if in == nil {
		return nil
	}
	out := new(MembershipBindingState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembershipState) DeepCopyInto(out *MembershipState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembershipState.
func (in *MembershipState) DeepCopy() *MembershipState {
	if in == nil {
		return nil
	}
	out := new(MembershipState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeState) DeepCopyInto(out *ScopeState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeState.
func (in *ScopeState) DeepCopy() *ScopeState {
	if in == nil {
		return nil
	}
	out := new(ScopeState)
	in.DeepCopyInto(out)
	return out
}
