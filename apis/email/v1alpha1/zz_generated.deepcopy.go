//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailIntegration) DeepCopyInto(out *EmailIntegration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailIntegration.
func (in *EmailIntegration) DeepCopy() *EmailIntegration {
	if in == nil {
		return nil
	}
	out := new(EmailIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailIntegration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailIntegrationInitParameters) DeepCopyInto(out *EmailIntegrationInitParameters) {
	*out = *in
	if in.EmailUsername != nil {
		in, out := &in.EmailUsername, &out.EmailUsername
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreRespondersFromPayload != nil {
		in, out := &in.IgnoreRespondersFromPayload, &out.IgnoreRespondersFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerTeamID != nil {
		in, out := &in.OwnerTeamID, &out.OwnerTeamID
		*out = new(string)
		**out = **in
	}
	if in.OwnerTeamIDRef != nil {
		in, out := &in.OwnerTeamIDRef, &out.OwnerTeamIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.OwnerTeamIDSelector != nil {
		in, out := &in.OwnerTeamIDSelector, &out.OwnerTeamIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]RespondersInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuppressNotifications != nil {
		in, out := &in.SuppressNotifications, &out.SuppressNotifications
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailIntegrationInitParameters.
func (in *EmailIntegrationInitParameters) DeepCopy() *EmailIntegrationInitParameters {
	if in == nil {
		return nil
	}
	out := new(EmailIntegrationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailIntegrationList) DeepCopyInto(out *EmailIntegrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmailIntegration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailIntegrationList.
func (in *EmailIntegrationList) DeepCopy() *EmailIntegrationList {
	if in == nil {
		return nil
	}
	out := new(EmailIntegrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailIntegrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailIntegrationObservation) DeepCopyInto(out *EmailIntegrationObservation) {
	*out = *in
	if in.EmailUsername != nil {
		in, out := &in.EmailUsername, &out.EmailUsername
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IgnoreRespondersFromPayload != nil {
		in, out := &in.IgnoreRespondersFromPayload, &out.IgnoreRespondersFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerTeamID != nil {
		in, out := &in.OwnerTeamID, &out.OwnerTeamID
		*out = new(string)
		**out = **in
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]RespondersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuppressNotifications != nil {
		in, out := &in.SuppressNotifications, &out.SuppressNotifications
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailIntegrationObservation.
func (in *EmailIntegrationObservation) DeepCopy() *EmailIntegrationObservation {
	if in == nil {
		return nil
	}
	out := new(EmailIntegrationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailIntegrationParameters) DeepCopyInto(out *EmailIntegrationParameters) {
	*out = *in
	if in.EmailUsername != nil {
		in, out := &in.EmailUsername, &out.EmailUsername
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreRespondersFromPayload != nil {
		in, out := &in.IgnoreRespondersFromPayload, &out.IgnoreRespondersFromPayload
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerTeamID != nil {
		in, out := &in.OwnerTeamID, &out.OwnerTeamID
		*out = new(string)
		**out = **in
	}
	if in.OwnerTeamIDRef != nil {
		in, out := &in.OwnerTeamIDRef, &out.OwnerTeamIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.OwnerTeamIDSelector != nil {
		in, out := &in.OwnerTeamIDSelector, &out.OwnerTeamIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]RespondersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuppressNotifications != nil {
		in, out := &in.SuppressNotifications, &out.SuppressNotifications
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailIntegrationParameters.
func (in *EmailIntegrationParameters) DeepCopy() *EmailIntegrationParameters {
	if in == nil {
		return nil
	}
	out := new(EmailIntegrationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailIntegrationSpec) DeepCopyInto(out *EmailIntegrationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailIntegrationSpec.
func (in *EmailIntegrationSpec) DeepCopy() *EmailIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(EmailIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailIntegrationStatus) DeepCopyInto(out *EmailIntegrationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailIntegrationStatus.
func (in *EmailIntegrationStatus) DeepCopy() *EmailIntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(EmailIntegrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RespondersInitParameters) DeepCopyInto(out *RespondersInitParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IDRef != nil {
		in, out := &in.IDRef, &out.IDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IDSelector != nil {
		in, out := &in.IDSelector, &out.IDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RespondersInitParameters.
func (in *RespondersInitParameters) DeepCopy() *RespondersInitParameters {
	if in == nil {
		return nil
	}
	out := new(RespondersInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RespondersObservation) DeepCopyInto(out *RespondersObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RespondersObservation.
func (in *RespondersObservation) DeepCopy() *RespondersObservation {
	if in == nil {
		return nil
	}
	out := new(RespondersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RespondersParameters) DeepCopyInto(out *RespondersParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IDRef != nil {
		in, out := &in.IDRef, &out.IDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IDSelector != nil {
		in, out := &in.IDSelector, &out.IDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RespondersParameters.
func (in *RespondersParameters) DeepCopy() *RespondersParameters {
	if in == nil {
		return nil
	}
	out := new(RespondersParameters)
	in.DeepCopyInto(out)
	return out
}
