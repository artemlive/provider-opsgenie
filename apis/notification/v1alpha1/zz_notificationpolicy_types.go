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

type AutoCloseActionInitParameters struct {

	// Duration of this action. This is a block, structure is documented below.
	Duration []DurationInitParameters `json:"duration,omitempty" tf:"duration,omitempty"`
}

type AutoCloseActionObservation struct {

	// Duration of this action. This is a block, structure is documented below.
	Duration []DurationObservation `json:"duration,omitempty" tf:"duration,omitempty"`
}

type AutoCloseActionParameters struct {

	// Duration of this action. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Duration []DurationParameters `json:"duration" tf:"duration,omitempty"`
}

type AutoRestartActionDurationInitParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type AutoRestartActionDurationObservation struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type AutoRestartActionDurationParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	// +kubebuilder:validation:Optional
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type AutoRestartActionInitParameters struct {

	// Duration of this action. This is a block, structure is documented below.
	Duration []AutoRestartActionDurationInitParameters `json:"duration,omitempty" tf:"duration,omitempty"`

	// How many times to repeat. This is a integer attribute.
	MaxRepeatCount *float64 `json:"maxRepeatCount,omitempty" tf:"max_repeat_count,omitempty"`
}

type AutoRestartActionObservation struct {

	// Duration of this action. This is a block, structure is documented below.
	Duration []AutoRestartActionDurationObservation `json:"duration,omitempty" tf:"duration,omitempty"`

	// How many times to repeat. This is a integer attribute.
	MaxRepeatCount *float64 `json:"maxRepeatCount,omitempty" tf:"max_repeat_count,omitempty"`
}

type AutoRestartActionParameters struct {

	// Duration of this action. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Duration []AutoRestartActionDurationParameters `json:"duration" tf:"duration,omitempty"`

	// How many times to repeat. This is a integer attribute.
	// +kubebuilder:validation:Optional
	MaxRepeatCount *float64 `json:"maxRepeatCount" tf:"max_repeat_count,omitempty"`
}

type ConditionsInitParameters struct {

	// User defined value that will be compared with alert field according to the operation. Default: empty string
	// User defined value that will be compared with alert field according to the operation. Default value is empty string
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Specifies which alert field will be used in condition. Possible values are message, alias, description, source, entity, tags, actions, details, extra-properties, responders, teams, priority
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// If field is set as extra-properties, key could be used for key-value pair
	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default: false
	// Indicates behaviour of the given operation. Default value is false
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// It is the operation that will be executed for the given field and key. Possible operations are matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty, equals-ignore-whitespace.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	// Order of the condition in conditions list
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type ConditionsObservation struct {

	// User defined value that will be compared with alert field according to the operation. Default: empty string
	// User defined value that will be compared with alert field according to the operation. Default value is empty string
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Specifies which alert field will be used in condition. Possible values are message, alias, description, source, entity, tags, actions, details, extra-properties, responders, teams, priority
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// If field is set as extra-properties, key could be used for key-value pair
	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default: false
	// Indicates behaviour of the given operation. Default value is false
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// It is the operation that will be executed for the given field and key. Possible operations are matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty, equals-ignore-whitespace.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	// Order of the condition in conditions list
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type ConditionsParameters struct {

	// User defined value that will be compared with alert field according to the operation. Default: empty string
	// User defined value that will be compared with alert field according to the operation. Default value is empty string
	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Specifies which alert field will be used in condition. Possible values are message, alias, description, source, entity, tags, actions, details, extra-properties, responders, teams, priority
	// +kubebuilder:validation:Optional
	Field *string `json:"field" tf:"field,omitempty"`

	// If field is set as extra-properties, key could be used for key-value pair
	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default: false
	// Indicates behaviour of the given operation. Default value is false
	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// It is the operation that will be executed for the given field and key. Possible operations are matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty, equals-ignore-whitespace.
	// +kubebuilder:validation:Optional
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	// Order of the condition in conditions list
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type DeDuplicationActionDurationInitParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DeDuplicationActionDurationObservation struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DeDuplicationActionDurationParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	// +kubebuilder:validation:Optional
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DeDuplicationActionInitParameters struct {

	// - Count
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Deduplication type. Possible values are: "value-based", "frequency-based"
	DeDuplicationActionType *string `json:"deDuplicationActionType,omitempty" tf:"de_duplication_action_type,omitempty"`

	// Duration of this action (only required for "frequency-based" de-duplication action). This is a block, structure is documented below.
	// Required when `de_duplication_action_type = "frequency-based"`
	Duration []DeDuplicationActionDurationInitParameters `json:"duration,omitempty" tf:"duration,omitempty"`
}

type DeDuplicationActionObservation struct {

	// - Count
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Deduplication type. Possible values are: "value-based", "frequency-based"
	DeDuplicationActionType *string `json:"deDuplicationActionType,omitempty" tf:"de_duplication_action_type,omitempty"`

	// Duration of this action (only required for "frequency-based" de-duplication action). This is a block, structure is documented below.
	// Required when `de_duplication_action_type = "frequency-based"`
	Duration []DeDuplicationActionDurationObservation `json:"duration,omitempty" tf:"duration,omitempty"`
}

type DeDuplicationActionParameters struct {

	// - Count
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// Deduplication type. Possible values are: "value-based", "frequency-based"
	// +kubebuilder:validation:Optional
	DeDuplicationActionType *string `json:"deDuplicationActionType" tf:"de_duplication_action_type,omitempty"`

	// Duration of this action (only required for "frequency-based" de-duplication action). This is a block, structure is documented below.
	// Required when `de_duplication_action_type = "frequency-based"`
	// +kubebuilder:validation:Optional
	Duration []DeDuplicationActionDurationParameters `json:"duration,omitempty" tf:"duration,omitempty"`
}

type DelayActionDurationInitParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DelayActionDurationObservation struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DelayActionDurationParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	// +kubebuilder:validation:Optional
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DelayActionInitParameters struct {

	// Defines until what day to delay or for what duration. Possible values are: for-duration, next-time, next-weekday, next-monday, next-tuesday, next-wednesday, next-thursday, next-friday, next-saturday, next-sunday
	DelayOption *string `json:"delayOption,omitempty" tf:"delay_option,omitempty"`

	// Duration of this action. If delay_option = for-duration this has to be set. This is a block, structure is documented below.
	// Required when `delay_option = "for-duration"`
	Duration []DelayActionDurationInitParameters `json:"duration,omitempty" tf:"duration,omitempty"`

	// Until what hour notifications will be delayed. If delay_option is set to antyhing else then for-duration this has to be set.
	UntilHour *float64 `json:"untilHour,omitempty" tf:"until_hour,omitempty"`

	// Until what minute on until_hour notifications will be delayed. If delay_option is set to antyhing else then for-duration this has to be set.
	UntilMinute *float64 `json:"untilMinute,omitempty" tf:"until_minute,omitempty"`
}

type DelayActionObservation struct {

	// Defines until what day to delay or for what duration. Possible values are: for-duration, next-time, next-weekday, next-monday, next-tuesday, next-wednesday, next-thursday, next-friday, next-saturday, next-sunday
	DelayOption *string `json:"delayOption,omitempty" tf:"delay_option,omitempty"`

	// Duration of this action. If delay_option = for-duration this has to be set. This is a block, structure is documented below.
	// Required when `delay_option = "for-duration"`
	Duration []DelayActionDurationObservation `json:"duration,omitempty" tf:"duration,omitempty"`

	// Until what hour notifications will be delayed. If delay_option is set to antyhing else then for-duration this has to be set.
	UntilHour *float64 `json:"untilHour,omitempty" tf:"until_hour,omitempty"`

	// Until what minute on until_hour notifications will be delayed. If delay_option is set to antyhing else then for-duration this has to be set.
	UntilMinute *float64 `json:"untilMinute,omitempty" tf:"until_minute,omitempty"`
}

type DelayActionParameters struct {

	// Defines until what day to delay or for what duration. Possible values are: for-duration, next-time, next-weekday, next-monday, next-tuesday, next-wednesday, next-thursday, next-friday, next-saturday, next-sunday
	// +kubebuilder:validation:Optional
	DelayOption *string `json:"delayOption" tf:"delay_option,omitempty"`

	// Duration of this action. If delay_option = for-duration this has to be set. This is a block, structure is documented below.
	// Required when `delay_option = "for-duration"`
	// +kubebuilder:validation:Optional
	Duration []DelayActionDurationParameters `json:"duration,omitempty" tf:"duration,omitempty"`

	// Until what hour notifications will be delayed. If delay_option is set to antyhing else then for-duration this has to be set.
	// +kubebuilder:validation:Optional
	UntilHour *float64 `json:"untilHour,omitempty" tf:"until_hour,omitempty"`

	// Until what minute on until_hour notifications will be delayed. If delay_option is set to antyhing else then for-duration this has to be set.
	// +kubebuilder:validation:Optional
	UntilMinute *float64 `json:"untilMinute,omitempty" tf:"until_minute,omitempty"`
}

type DurationInitParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DurationObservation struct {

	// A amount of time in time_units. This is a integer attribute.
	TimeAmount *float64 `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DurationParameters struct {

	// A amount of time in time_units. This is a integer attribute.
	// +kubebuilder:validation:Optional
	TimeAmount *float64 `json:"timeAmount" tf:"time_amount,omitempty"`

	// Valid time units are: minutes, hours, days. Default: minutes
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type FilterInitParameters struct {

	// Conditions applied to filter. This is a block, structure is documented below.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// A filter type, supported types are: match-all, match-any-condition, match-all-conditions. Default: match-all
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FilterObservation struct {

	// Conditions applied to filter. This is a block, structure is documented below.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// A filter type, supported types are: match-all, match-any-condition, match-all-conditions. Default: match-all
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FilterParameters struct {

	// Conditions applied to filter. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// A filter type, supported types are: match-all, match-any-condition, match-all-conditions. Default: match-all
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NotificationPolicyInitParameters struct {

	// Auto Restart Action of the policy. This is a block, structure is documented below.
	AutoCloseAction []AutoCloseActionInitParameters `json:"autoCloseAction,omitempty" tf:"auto_close_action,omitempty"`

	// Auto Restart Action of the policy. This is a block, structure is documented below.
	AutoRestartAction []AutoRestartActionInitParameters `json:"autoRestartAction,omitempty" tf:"auto_restart_action,omitempty"`

	// Deduplication Action of the policy. This is a block, structure is documented below.
	DeDuplicationAction []DeDuplicationActionInitParameters `json:"deDuplicationAction,omitempty" tf:"de_duplication_action,omitempty"`

	// Delay notifications. This is a block, structure is documented below.
	DelayAction []DelayActionInitParameters `json:"delayAction,omitempty" tf:"delay_action,omitempty"`

	// If policy should be enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A notification filter which will be applied. This filter can be empty: filter {} - this means match-all. This is a block, structure is documented below.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Description of the policy. This can be max 512 characters.
	PolicyDescription *string `json:"policyDescription,omitempty" tf:"policy_description,omitempty"`

	// Suppress value of the policy. Values are: true, false. Default: false
	Suppress *bool `json:"suppress,omitempty" tf:"suppress,omitempty"`

	// Id of team that this policy belons to.
	// +crossplane:generate:reference:type=github.com/macpaw/provider-opsgenie/apis/team/v1alpha1.Team
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`

	// Time restrictions specified in this field must be met for this policy to work. This is a block, structure is documented below.
	TimeRestriction []TimeRestrictionInitParameters `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`
}

type NotificationPolicyObservation struct {

	// Auto Restart Action of the policy. This is a block, structure is documented below.
	AutoCloseAction []AutoCloseActionObservation `json:"autoCloseAction,omitempty" tf:"auto_close_action,omitempty"`

	// Auto Restart Action of the policy. This is a block, structure is documented below.
	AutoRestartAction []AutoRestartActionObservation `json:"autoRestartAction,omitempty" tf:"auto_restart_action,omitempty"`

	// Deduplication Action of the policy. This is a block, structure is documented below.
	DeDuplicationAction []DeDuplicationActionObservation `json:"deDuplicationAction,omitempty" tf:"de_duplication_action,omitempty"`

	// Delay notifications. This is a block, structure is documented below.
	DelayAction []DelayActionObservation `json:"delayAction,omitempty" tf:"delay_action,omitempty"`

	// If policy should be enabled. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A notification filter which will be applied. This filter can be empty: filter {} - this means match-all. This is a block, structure is documented below.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// The ID of the Opsgenie Notification Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Description of the policy. This can be max 512 characters.
	PolicyDescription *string `json:"policyDescription,omitempty" tf:"policy_description,omitempty"`

	// Suppress value of the policy. Values are: true, false. Default: false
	Suppress *bool `json:"suppress,omitempty" tf:"suppress,omitempty"`

	// Id of team that this policy belons to.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Time restrictions specified in this field must be met for this policy to work. This is a block, structure is documented below.
	TimeRestriction []TimeRestrictionObservation `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`
}

type NotificationPolicyParameters struct {

	// Auto Restart Action of the policy. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	AutoCloseAction []AutoCloseActionParameters `json:"autoCloseAction,omitempty" tf:"auto_close_action,omitempty"`

	// Auto Restart Action of the policy. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	AutoRestartAction []AutoRestartActionParameters `json:"autoRestartAction,omitempty" tf:"auto_restart_action,omitempty"`

	// Deduplication Action of the policy. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	DeDuplicationAction []DeDuplicationActionParameters `json:"deDuplicationAction,omitempty" tf:"de_duplication_action,omitempty"`

	// Delay notifications. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	DelayAction []DelayActionParameters `json:"delayAction,omitempty" tf:"delay_action,omitempty"`

	// If policy should be enabled. Default: true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A notification filter which will be applied. This filter can be empty: filter {} - this means match-all. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Description of the policy. This can be max 512 characters.
	// +kubebuilder:validation:Optional
	PolicyDescription *string `json:"policyDescription,omitempty" tf:"policy_description,omitempty"`

	// Suppress value of the policy. Values are: true, false. Default: false
	// +kubebuilder:validation:Optional
	Suppress *bool `json:"suppress,omitempty" tf:"suppress,omitempty"`

	// Id of team that this policy belons to.
	// +crossplane:generate:reference:type=github.com/macpaw/provider-opsgenie/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`

	// Time restrictions specified in this field must be met for this policy to work. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	TimeRestriction []TimeRestrictionParameters `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`
}

type RestrictionInitParameters struct {

	// Ending hour of restriction on defined end_day
	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	// Starting hour of restriction on defined start_day
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type RestrictionObservation struct {

	// Ending hour of restriction on defined end_day
	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	// Starting hour of restriction on defined start_day
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type RestrictionParameters struct {

	// Ending hour of restriction on defined end_day
	// +kubebuilder:validation:Optional
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	// +kubebuilder:validation:Optional
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// Starting hour of restriction on defined start_day
	// +kubebuilder:validation:Optional
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	// +kubebuilder:validation:Optional
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type RestrictionsInitParameters struct {

	// Ending day of restriction (eg. wednesday)
	EndDay *string `json:"endDay,omitempty" tf:"end_day,omitempty"`

	// Ending hour of restriction on defined end_day
	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	// Starting day of restriction (eg. monday)
	StartDay *string `json:"startDay,omitempty" tf:"start_day,omitempty"`

	// Starting hour of restriction on defined start_day
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type RestrictionsObservation struct {

	// Ending day of restriction (eg. wednesday)
	EndDay *string `json:"endDay,omitempty" tf:"end_day,omitempty"`

	// Ending hour of restriction on defined end_day
	EndHour *float64 `json:"endHour,omitempty" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	EndMin *float64 `json:"endMin,omitempty" tf:"end_min,omitempty"`

	// Starting day of restriction (eg. monday)
	StartDay *string `json:"startDay,omitempty" tf:"start_day,omitempty"`

	// Starting hour of restriction on defined start_day
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	StartMin *float64 `json:"startMin,omitempty" tf:"start_min,omitempty"`
}

type RestrictionsParameters struct {

	// Ending day of restriction (eg. wednesday)
	// +kubebuilder:validation:Optional
	EndDay *string `json:"endDay" tf:"end_day,omitempty"`

	// Ending hour of restriction on defined end_day
	// +kubebuilder:validation:Optional
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	// +kubebuilder:validation:Optional
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// Starting day of restriction (eg. monday)
	// +kubebuilder:validation:Optional
	StartDay *string `json:"startDay" tf:"start_day,omitempty"`

	// Starting hour of restriction on defined start_day
	// +kubebuilder:validation:Optional
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	// +kubebuilder:validation:Optional
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type TimeRestrictionInitParameters struct {

	// A definition of hourly definition applied daily, this has to be used with combination: type = time-of-day. This is a block, structure is documented below.
	Restriction []RestrictionInitParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// List of days and hours definitions for field type = weekday-and-time-of-day. This is a block, structure is documented below.
	Restrictions []RestrictionsInitParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// Defines if restriction should apply daily on given hours or on certain days and hours. Possible values are: time-of-day, weekday-and-time-of-day
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TimeRestrictionObservation struct {

	// A definition of hourly definition applied daily, this has to be used with combination: type = time-of-day. This is a block, structure is documented below.
	Restriction []RestrictionObservation `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// List of days and hours definitions for field type = weekday-and-time-of-day. This is a block, structure is documented below.
	Restrictions []RestrictionsObservation `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// Defines if restriction should apply daily on given hours or on certain days and hours. Possible values are: time-of-day, weekday-and-time-of-day
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TimeRestrictionParameters struct {

	// A definition of hourly definition applied daily, this has to be used with combination: type = time-of-day. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Restriction []RestrictionParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// List of days and hours definitions for field type = weekday-and-time-of-day. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Restrictions []RestrictionsParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// Defines if restriction should apply daily on given hours or on certain days and hours. Possible values are: time-of-day, weekday-and-time-of-day
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// NotificationPolicySpec defines the desired state of NotificationPolicy
type NotificationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationPolicyParameters `json:"forProvider"`
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
	InitProvider NotificationPolicyInitParameters `json:"initProvider,omitempty"`
}

// NotificationPolicyStatus defines the observed state of NotificationPolicy.
type NotificationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NotificationPolicy is the Schema for the NotificationPolicys API. Manages a Notification Policy within Opsgenie.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie}
type NotificationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filter) || (has(self.initProvider) && has(self.initProvider.filter))",message="spec.forProvider.filter is a required parameter"
	Spec   NotificationPolicySpec   `json:"spec"`
	Status NotificationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationPolicyList contains a list of NotificationPolicys
type NotificationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationPolicy `json:"items"`
}

// Repository type metadata.
var (
	NotificationPolicy_Kind             = "NotificationPolicy"
	NotificationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationPolicy_Kind}.String()
	NotificationPolicy_KindAPIVersion   = NotificationPolicy_Kind + "." + CRDGroupVersion.String()
	NotificationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(NotificationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationPolicy{}, &NotificationPolicyList{})
}
