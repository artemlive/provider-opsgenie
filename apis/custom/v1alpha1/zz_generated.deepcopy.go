//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRole) DeepCopyInto(out *CustomRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRole.
func (in *CustomRole) DeepCopy() *CustomRole {
	if in == nil {
		return nil
	}
	out := new(CustomRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleInitParameters) DeepCopyInto(out *CustomRoleInitParameters) {
	*out = *in
	if in.DisallowedRights != nil {
		in, out := &in.DisallowedRights, &out.DisallowedRights
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExtendedRole != nil {
		in, out := &in.ExtendedRole, &out.ExtendedRole
		*out = new(string)
		**out = **in
	}
	if in.GrantedRights != nil {
		in, out := &in.GrantedRights, &out.GrantedRights
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleInitParameters.
func (in *CustomRoleInitParameters) DeepCopy() *CustomRoleInitParameters {
	if in == nil {
		return nil
	}
	out := new(CustomRoleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleList) DeepCopyInto(out *CustomRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleList.
func (in *CustomRoleList) DeepCopy() *CustomRoleList {
	if in == nil {
		return nil
	}
	out := new(CustomRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleObservation) DeepCopyInto(out *CustomRoleObservation) {
	*out = *in
	if in.DisallowedRights != nil {
		in, out := &in.DisallowedRights, &out.DisallowedRights
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExtendedRole != nil {
		in, out := &in.ExtendedRole, &out.ExtendedRole
		*out = new(string)
		**out = **in
	}
	if in.GrantedRights != nil {
		in, out := &in.GrantedRights, &out.GrantedRights
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleObservation.
func (in *CustomRoleObservation) DeepCopy() *CustomRoleObservation {
	if in == nil {
		return nil
	}
	out := new(CustomRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleParameters) DeepCopyInto(out *CustomRoleParameters) {
	*out = *in
	if in.DisallowedRights != nil {
		in, out := &in.DisallowedRights, &out.DisallowedRights
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExtendedRole != nil {
		in, out := &in.ExtendedRole, &out.ExtendedRole
		*out = new(string)
		**out = **in
	}
	if in.GrantedRights != nil {
		in, out := &in.GrantedRights, &out.GrantedRights
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleParameters.
func (in *CustomRoleParameters) DeepCopy() *CustomRoleParameters {
	if in == nil {
		return nil
	}
	out := new(CustomRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleSpec) DeepCopyInto(out *CustomRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleSpec.
func (in *CustomRoleSpec) DeepCopy() *CustomRoleSpec {
	if in == nil {
		return nil
	}
	out := new(CustomRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRoleStatus) DeepCopyInto(out *CustomRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRoleStatus.
func (in *CustomRoleStatus) DeepCopy() *CustomRoleStatus {
	if in == nil {
		return nil
	}
	out := new(CustomRoleStatus)
	in.DeepCopyInto(out)
	return out
}
