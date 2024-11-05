// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ContactInitParameters struct {

	// Contact method. Possible values: email, sms, voice, mobile
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Address of a given method (eg. email address for email, phone number for sms/voice or mobile application name for mobile)
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type ContactObservation struct {

	// Contact method. Possible values: email, sms, voice, mobile
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Address of a given method (eg. email address for email, phone number for sms/voice or mobile application name for mobile)
	To *string `json:"to,omitempty" tf:"to,omitempty"`
}

type ContactParameters struct {

	// Contact method. Possible values: email, sms, voice, mobile
	// +kubebuilder:validation:Optional
	Method *string `json:"method" tf:"method,omitempty"`

	// Address of a given method (eg. email address for email, phone number for sms/voice or mobile application name for mobile)
	// +kubebuilder:validation:Optional
	To *string `json:"to" tf:"to,omitempty"`
}

type CriteriaConditionsInitParameters struct {

	// User defined value that will be compared with alert field according to the operation. Default: empty string
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Possible values: message, alias, description, source, entity, tags, actions, details, extra-properties, recipients, teams, priority
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default: false
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// Possible values: matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty, equals-ignore-whitespace
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type CriteriaConditionsObservation struct {

	// User defined value that will be compared with alert field according to the operation. Default: empty string
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Possible values: message, alias, description, source, entity, tags, actions, details, extra-properties, recipients, teams, priority
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default: false
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// Possible values: matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty, equals-ignore-whitespace
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type CriteriaConditionsParameters struct {

	// User defined value that will be compared with alert field according to the operation. Default: empty string
	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Possible values: message, alias, description, source, entity, tags, actions, details, extra-properties, recipients, teams, priority
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default: false
	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// Possible values: matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty, equals-ignore-whitespace
	// +kubebuilder:validation:Optional
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type CriteriaInitParameters struct {

	// Defines the fields and values when the condition applies
	Conditions []CriteriaConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CriteriaObservation struct {

	// Defines the fields and values when the condition applies
	Conditions []CriteriaConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CriteriaParameters struct {

	// Defines the fields and values when the condition applies
	// +kubebuilder:validation:Optional
	Conditions []CriteriaConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NotificationRuleInitParameters struct {

	// Type of the action that notification rule will have. Allowed values: create-alert, acknowledged-alert, closed-alert, assigned-alert, add-note, schedule-start, schedule-end, incoming-call-routing
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	Criteria []CriteriaInitParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// If policy should be enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of Time Periods that notification for schedule start/end will be sent. Allowed values: just-before, 15-minutes-ago, 1-hour-ago, 1-day-ago. If action_type is schedule-start or schedule-end then it is required.
	// +listType=set
	NotificationTime []*string `json:"notificationTime,omitempty" tf:"notification_time,omitempty"`

	// Order of the condition in conditions list
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	Repeat []RepeatInitParameters `json:"repeat,omitempty" tf:"repeat,omitempty"`

	Schedules []SchedulesInitParameters `json:"schedules,omitempty" tf:"schedules,omitempty"`

	// Notification rule steps to take (eg. SMS or email message). This is a block, structure is documented below.
	Steps []StepsInitParameters `json:"steps,omitempty" tf:"steps,omitempty"`

	TimeRestriction []NotificationRuleTimeRestrictionInitParameters `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`

	// Username of user to which this notification rule belongs to.
	// +crossplane:generate:reference:type=github.com/macpaw/provider-opsgenie/apis/user/v1alpha1.User
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User in user to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User in user to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

type NotificationRuleObservation struct {

	// Type of the action that notification rule will have. Allowed values: create-alert, acknowledged-alert, closed-alert, assigned-alert, add-note, schedule-start, schedule-end, incoming-call-routing
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	Criteria []CriteriaObservation `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// If policy should be enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Opsgenie Notification Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of Time Periods that notification for schedule start/end will be sent. Allowed values: just-before, 15-minutes-ago, 1-hour-ago, 1-day-ago. If action_type is schedule-start or schedule-end then it is required.
	// +listType=set
	NotificationTime []*string `json:"notificationTime,omitempty" tf:"notification_time,omitempty"`

	// Order of the condition in conditions list
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	Repeat []RepeatObservation `json:"repeat,omitempty" tf:"repeat,omitempty"`

	Schedules []SchedulesObservation `json:"schedules,omitempty" tf:"schedules,omitempty"`

	// Notification rule steps to take (eg. SMS or email message). This is a block, structure is documented below.
	Steps []StepsObservation `json:"steps,omitempty" tf:"steps,omitempty"`

	TimeRestriction []NotificationRuleTimeRestrictionObservation `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`

	// Username of user to which this notification rule belongs to.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type NotificationRuleParameters struct {

	// Type of the action that notification rule will have. Allowed values: create-alert, acknowledged-alert, closed-alert, assigned-alert, add-note, schedule-start, schedule-end, incoming-call-routing
	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Criteria []CriteriaParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// If policy should be enabled. Default: true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of Time Periods that notification for schedule start/end will be sent. Allowed values: just-before, 15-minutes-ago, 1-hour-ago, 1-day-ago. If action_type is schedule-start or schedule-end then it is required.
	// +kubebuilder:validation:Optional
	// +listType=set
	NotificationTime []*string `json:"notificationTime,omitempty" tf:"notification_time,omitempty"`

	// Order of the condition in conditions list
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// +kubebuilder:validation:Optional
	Repeat []RepeatParameters `json:"repeat,omitempty" tf:"repeat,omitempty"`

	// +kubebuilder:validation:Optional
	Schedules []SchedulesParameters `json:"schedules,omitempty" tf:"schedules,omitempty"`

	// Notification rule steps to take (eg. SMS or email message). This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Steps []StepsParameters `json:"steps,omitempty" tf:"steps,omitempty"`

	// +kubebuilder:validation:Optional
	TimeRestriction []NotificationRuleTimeRestrictionParameters `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`

	// Username of user to which this notification rule belongs to.
	// +crossplane:generate:reference:type=github.com/macpaw/provider-opsgenie/apis/user/v1alpha1.User
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User in user to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User in user to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

type NotificationRuleTimeRestrictionInitParameters struct {
	Restriction []TimeRestrictionRestrictionInitParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	Restrictions []TimeRestrictionRestrictionsInitParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NotificationRuleTimeRestrictionObservation struct {
	Restriction []TimeRestrictionRestrictionObservation `json:"restriction,omitempty" tf:"restriction,omitempty"`

	Restrictions []TimeRestrictionRestrictionsObservation `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NotificationRuleTimeRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	Restriction []TimeRestrictionRestrictionParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// +kubebuilder:validation:Optional
	Restrictions []TimeRestrictionRestrictionsParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RepeatInitParameters struct {

	// Defined if this step is enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	LoopAfter *float64 `json:"loopAfter,omitempty" tf:"loop_after,omitempty"`
}

type RepeatObservation struct {

	// Defined if this step is enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	LoopAfter *float64 `json:"loopAfter,omitempty" tf:"loop_after,omitempty"`
}

type RepeatParameters struct {

	// Defined if this step is enabled. Default: true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LoopAfter *float64 `json:"loopAfter" tf:"loop_after,omitempty"`
}

type SchedulesInitParameters struct {

	// Name of the notification policy
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchedulesObservation struct {

	// Name of the notification policy
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchedulesParameters struct {

	// Name of the notification policy
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Kind of matching filter. Possible values: match-all, match-any-condition, match-all-conditions
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type StepsInitParameters struct {

	// Defines the contact that notification will be sent to. This is a block, structure is documented below.
	Contact []ContactInitParameters `json:"contact,omitempty" tf:"contact,omitempty"`

	// Defined if this step is enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Time period, in minutes, notification will be sent after.
	SendAfter *float64 `json:"sendAfter,omitempty" tf:"send_after,omitempty"`
}

type StepsObservation struct {

	// Defines the contact that notification will be sent to. This is a block, structure is documented below.
	Contact []ContactObservation `json:"contact,omitempty" tf:"contact,omitempty"`

	// Defined if this step is enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Time period, in minutes, notification will be sent after.
	SendAfter *float64 `json:"sendAfter,omitempty" tf:"send_after,omitempty"`
}

type StepsParameters struct {

	// Defines the contact that notification will be sent to. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Contact []ContactParameters `json:"contact" tf:"contact,omitempty"`

	// Defined if this step is enabled. Default: true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Time period, in minutes, notification will be sent after.
	// +kubebuilder:validation:Optional
	SendAfter *float64 `json:"sendAfter,omitempty" tf:"send_after,omitempty"`
}

type TimeRestrictionRestrictionInitParameters struct {
	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type TimeRestrictionRestrictionObservation struct {
	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type TimeRestrictionRestrictionParameters struct {

	// +kubebuilder:validation:Optional
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// +kubebuilder:validation:Optional
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// +kubebuilder:validation:Optional
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Optional
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type TimeRestrictionRestrictionsInitParameters struct {
	EndDay *string `json:"endDay,omitempty" tf:"end_day,omitempty"`

	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	StartDay *string `json:"startDay,omitempty" tf:"start_day,omitempty"`

	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type TimeRestrictionRestrictionsObservation struct {
	EndDay *string `json:"endDay,omitempty" tf:"end_day,omitempty"`

	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	StartDay *string `json:"startDay,omitempty" tf:"start_day,omitempty"`

	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type TimeRestrictionRestrictionsParameters struct {

	// +kubebuilder:validation:Optional
	EndDay *string `json:"endDay" tf:"end_day,omitempty"`

	// +kubebuilder:validation:Optional
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// +kubebuilder:validation:Optional
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// +kubebuilder:validation:Optional
	StartDay *string `json:"startDay" tf:"start_day,omitempty"`

	// +kubebuilder:validation:Optional
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Optional
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

// NotificationRuleSpec defines the desired state of NotificationRule
type NotificationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationRuleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NotificationRuleInitParameters `json:"initProvider,omitempty"`
}

// NotificationRuleStatus defines the observed state of NotificationRule.
type NotificationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NotificationRule is the Schema for the NotificationRules API. Manages a Notification Rule within Opsgenie.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie}
type NotificationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionType) || (has(self.initProvider) && has(self.initProvider.actionType))",message="spec.forProvider.actionType is a required parameter"
	Spec   NotificationRuleSpec   `json:"spec"`
	Status NotificationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRuleList contains a list of NotificationRules
type NotificationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationRule `json:"items"`
}

// Repository type metadata.
var (
	NotificationRule_Kind             = "NotificationRule"
	NotificationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationRule_Kind}.String()
	NotificationRule_KindAPIVersion   = NotificationRule_Kind + "." + CRDGroupVersion.String()
	NotificationRule_GroupVersionKind = CRDGroupVersion.WithKind(NotificationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationRule{}, &NotificationRuleList{})
}