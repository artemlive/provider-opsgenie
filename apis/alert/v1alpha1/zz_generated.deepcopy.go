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
func (in *AlertPolicy) DeepCopyInto(out *AlertPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicy.
func (in *AlertPolicy) DeepCopy() *AlertPolicy {
	if in == nil {
		return nil
	}
	out := new(AlertPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicyInitParameters) DeepCopyInto(out *AlertPolicyInitParameters) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlertDescription != nil {
		in, out := &in.AlertDescription, &out.AlertDescription
		*out = new(string)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.ContinuePolicy != nil {
		in, out := &in.ContinuePolicy, &out.ContinuePolicy
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Entity != nil {
		in, out := &in.Entity, &out.Entity
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]FilterInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreOriginalActions != nil {
		in, out := &in.IgnoreOriginalActions, &out.IgnoreOriginalActions
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalDetails != nil {
		in, out := &in.IgnoreOriginalDetails, &out.IgnoreOriginalDetails
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalResponders != nil {
		in, out := &in.IgnoreOriginalResponders, &out.IgnoreOriginalResponders
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalTags != nil {
		in, out := &in.IgnoreOriginalTags, &out.IgnoreOriginalTags
		*out = new(bool)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyDescription != nil {
		in, out := &in.PolicyDescription, &out.PolicyDescription
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]RespondersInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.TeamIDRef != nil {
		in, out := &in.TeamIDRef, &out.TeamIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TeamIDSelector != nil {
		in, out := &in.TeamIDSelector, &out.TeamIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeRestriction != nil {
		in, out := &in.TimeRestriction, &out.TimeRestriction
		*out = make([]TimeRestrictionInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicyInitParameters.
func (in *AlertPolicyInitParameters) DeepCopy() *AlertPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(AlertPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicyList) DeepCopyInto(out *AlertPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicyList.
func (in *AlertPolicyList) DeepCopy() *AlertPolicyList {
	if in == nil {
		return nil
	}
	out := new(AlertPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicyObservation) DeepCopyInto(out *AlertPolicyObservation) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlertDescription != nil {
		in, out := &in.AlertDescription, &out.AlertDescription
		*out = new(string)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.ContinuePolicy != nil {
		in, out := &in.ContinuePolicy, &out.ContinuePolicy
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Entity != nil {
		in, out := &in.Entity, &out.Entity
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]FilterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IgnoreOriginalActions != nil {
		in, out := &in.IgnoreOriginalActions, &out.IgnoreOriginalActions
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalDetails != nil {
		in, out := &in.IgnoreOriginalDetails, &out.IgnoreOriginalDetails
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalResponders != nil {
		in, out := &in.IgnoreOriginalResponders, &out.IgnoreOriginalResponders
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalTags != nil {
		in, out := &in.IgnoreOriginalTags, &out.IgnoreOriginalTags
		*out = new(bool)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyDescription != nil {
		in, out := &in.PolicyDescription, &out.PolicyDescription
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
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
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.TimeRestriction != nil {
		in, out := &in.TimeRestriction, &out.TimeRestriction
		*out = make([]TimeRestrictionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicyObservation.
func (in *AlertPolicyObservation) DeepCopy() *AlertPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(AlertPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicyParameters) DeepCopyInto(out *AlertPolicyParameters) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlertDescription != nil {
		in, out := &in.AlertDescription, &out.AlertDescription
		*out = new(string)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.ContinuePolicy != nil {
		in, out := &in.ContinuePolicy, &out.ContinuePolicy
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Entity != nil {
		in, out := &in.Entity, &out.Entity
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]FilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreOriginalActions != nil {
		in, out := &in.IgnoreOriginalActions, &out.IgnoreOriginalActions
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalDetails != nil {
		in, out := &in.IgnoreOriginalDetails, &out.IgnoreOriginalDetails
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalResponders != nil {
		in, out := &in.IgnoreOriginalResponders, &out.IgnoreOriginalResponders
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreOriginalTags != nil {
		in, out := &in.IgnoreOriginalTags, &out.IgnoreOriginalTags
		*out = new(bool)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyDescription != nil {
		in, out := &in.PolicyDescription, &out.PolicyDescription
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]RespondersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.TeamIDRef != nil {
		in, out := &in.TeamIDRef, &out.TeamIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TeamIDSelector != nil {
		in, out := &in.TeamIDSelector, &out.TeamIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeRestriction != nil {
		in, out := &in.TimeRestriction, &out.TimeRestriction
		*out = make([]TimeRestrictionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicyParameters.
func (in *AlertPolicyParameters) DeepCopy() *AlertPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(AlertPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicySpec) DeepCopyInto(out *AlertPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicySpec.
func (in *AlertPolicySpec) DeepCopy() *AlertPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AlertPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicyStatus) DeepCopyInto(out *AlertPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicyStatus.
func (in *AlertPolicyStatus) DeepCopy() *AlertPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AlertPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsInitParameters) DeepCopyInto(out *ConditionsInitParameters) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsInitParameters.
func (in *ConditionsInitParameters) DeepCopy() *ConditionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsObservation) DeepCopyInto(out *ConditionsObservation) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsObservation.
func (in *ConditionsObservation) DeepCopy() *ConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsParameters) DeepCopyInto(out *ConditionsParameters) {
	*out = *in
	if in.ExpectedValue != nil {
		in, out := &in.ExpectedValue, &out.ExpectedValue
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(bool)
		**out = **in
	}
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsParameters.
func (in *ConditionsParameters) DeepCopy() *ConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterInitParameters) DeepCopyInto(out *FilterInitParameters) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterInitParameters.
func (in *FilterInitParameters) DeepCopy() *FilterInitParameters {
	if in == nil {
		return nil
	}
	out := new(FilterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterObservation) DeepCopyInto(out *FilterObservation) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterObservation.
func (in *FilterObservation) DeepCopy() *FilterObservation {
	if in == nil {
		return nil
	}
	out := new(FilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterParameters) DeepCopyInto(out *FilterParameters) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterParameters.
func (in *FilterParameters) DeepCopy() *FilterParameters {
	if in == nil {
		return nil
	}
	out := new(FilterParameters)
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionInitParameters) DeepCopyInto(out *RestrictionInitParameters) {
	*out = *in
	if in.EndHour != nil {
		in, out := &in.EndHour, &out.EndHour
		*out = new(float64)
		**out = **in
	}
	if in.EndMin != nil {
		in, out := &in.EndMin, &out.EndMin
		*out = new(float64)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMin != nil {
		in, out := &in.StartMin, &out.StartMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionInitParameters.
func (in *RestrictionInitParameters) DeepCopy() *RestrictionInitParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionObservation) DeepCopyInto(out *RestrictionObservation) {
	*out = *in
	if in.EndHour != nil {
		in, out := &in.EndHour, &out.EndHour
		*out = new(float64)
		**out = **in
	}
	if in.EndMin != nil {
		in, out := &in.EndMin, &out.EndMin
		*out = new(float64)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMin != nil {
		in, out := &in.StartMin, &out.StartMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionObservation.
func (in *RestrictionObservation) DeepCopy() *RestrictionObservation {
	if in == nil {
		return nil
	}
	out := new(RestrictionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionParameters) DeepCopyInto(out *RestrictionParameters) {
	*out = *in
	if in.EndHour != nil {
		in, out := &in.EndHour, &out.EndHour
		*out = new(float64)
		**out = **in
	}
	if in.EndMin != nil {
		in, out := &in.EndMin, &out.EndMin
		*out = new(float64)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMin != nil {
		in, out := &in.StartMin, &out.StartMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionParameters.
func (in *RestrictionParameters) DeepCopy() *RestrictionParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionsInitParameters) DeepCopyInto(out *RestrictionsInitParameters) {
	*out = *in
	if in.EndDay != nil {
		in, out := &in.EndDay, &out.EndDay
		*out = new(string)
		**out = **in
	}
	if in.EndHour != nil {
		in, out := &in.EndHour, &out.EndHour
		*out = new(float64)
		**out = **in
	}
	if in.EndMin != nil {
		in, out := &in.EndMin, &out.EndMin
		*out = new(float64)
		**out = **in
	}
	if in.StartDay != nil {
		in, out := &in.StartDay, &out.StartDay
		*out = new(string)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMin != nil {
		in, out := &in.StartMin, &out.StartMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionsInitParameters.
func (in *RestrictionsInitParameters) DeepCopy() *RestrictionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionsObservation) DeepCopyInto(out *RestrictionsObservation) {
	*out = *in
	if in.EndDay != nil {
		in, out := &in.EndDay, &out.EndDay
		*out = new(string)
		**out = **in
	}
	if in.EndHour != nil {
		in, out := &in.EndHour, &out.EndHour
		*out = new(float64)
		**out = **in
	}
	if in.EndMin != nil {
		in, out := &in.EndMin, &out.EndMin
		*out = new(float64)
		**out = **in
	}
	if in.StartDay != nil {
		in, out := &in.StartDay, &out.StartDay
		*out = new(string)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMin != nil {
		in, out := &in.StartMin, &out.StartMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionsObservation.
func (in *RestrictionsObservation) DeepCopy() *RestrictionsObservation {
	if in == nil {
		return nil
	}
	out := new(RestrictionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestrictionsParameters) DeepCopyInto(out *RestrictionsParameters) {
	*out = *in
	if in.EndDay != nil {
		in, out := &in.EndDay, &out.EndDay
		*out = new(string)
		**out = **in
	}
	if in.EndHour != nil {
		in, out := &in.EndHour, &out.EndHour
		*out = new(float64)
		**out = **in
	}
	if in.EndMin != nil {
		in, out := &in.EndMin, &out.EndMin
		*out = new(float64)
		**out = **in
	}
	if in.StartDay != nil {
		in, out := &in.StartDay, &out.StartDay
		*out = new(string)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(float64)
		**out = **in
	}
	if in.StartMin != nil {
		in, out := &in.StartMin, &out.StartMin
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestrictionsParameters.
func (in *RestrictionsParameters) DeepCopy() *RestrictionsParameters {
	if in == nil {
		return nil
	}
	out := new(RestrictionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRestrictionInitParameters) DeepCopyInto(out *TimeRestrictionInitParameters) {
	*out = *in
	if in.Restriction != nil {
		in, out := &in.Restriction, &out.Restriction
		*out = make([]RestrictionInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restrictions != nil {
		in, out := &in.Restrictions, &out.Restrictions
		*out = make([]RestrictionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRestrictionInitParameters.
func (in *TimeRestrictionInitParameters) DeepCopy() *TimeRestrictionInitParameters {
	if in == nil {
		return nil
	}
	out := new(TimeRestrictionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRestrictionObservation) DeepCopyInto(out *TimeRestrictionObservation) {
	*out = *in
	if in.Restriction != nil {
		in, out := &in.Restriction, &out.Restriction
		*out = make([]RestrictionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restrictions != nil {
		in, out := &in.Restrictions, &out.Restrictions
		*out = make([]RestrictionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRestrictionObservation.
func (in *TimeRestrictionObservation) DeepCopy() *TimeRestrictionObservation {
	if in == nil {
		return nil
	}
	out := new(TimeRestrictionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRestrictionParameters) DeepCopyInto(out *TimeRestrictionParameters) {
	*out = *in
	if in.Restriction != nil {
		in, out := &in.Restriction, &out.Restriction
		*out = make([]RestrictionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restrictions != nil {
		in, out := &in.Restrictions, &out.Restrictions
		*out = make([]RestrictionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRestrictionParameters.
func (in *TimeRestrictionParameters) DeepCopy() *TimeRestrictionParameters {
	if in == nil {
		return nil
	}
	out := new(TimeRestrictionParameters)
	in.DeepCopyInto(out)
	return out
}