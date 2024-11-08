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

type IntegrationInitParameters struct {
	AllowConfigurationAccess *bool `json:"allowConfigurationAccess,omitempty" tf:"allow_configuration_access,omitempty"`

	// This parameter is for configuring the write access of integration. If write access is restricted, the integration will not be authorized to write within any domain. Default: true.
	AllowWriteAccess *bool `json:"allowWriteAccess,omitempty" tf:"allow_write_access,omitempty"`

	// This parameter is for specifying whether the integration will be enabled or not. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// If enabled, the integration will ignore recipients sent in request payloads. Default: false.
	IgnoreRespondersFromPayload *bool `json:"ignoreRespondersFromPayload,omitempty" tf:"ignore_responders_from_payload,omitempty"`

	// Name of the integration. Name must be unique for each integration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner team id of the integration. If changed, this will recreate a new API integration, which will probably have a different API key.
	OwnerTeamID *string `json:"ownerTeamId,omitempty" tf:"owner_team_id,omitempty"`

	// User, schedule, teams or escalation names to calculate which users will receive the notifications of the alert.
	Responders []RespondersInitParameters `json:"responders,omitempty" tf:"responders,omitempty"`

	// If enabled, notifications that come from alerts will be suppressed. Default: false.
	SuppressNotifications *bool `json:"suppressNotifications,omitempty" tf:"suppress_notifications,omitempty"`

	// Type of the integration (API, Marid, Prometheus, etc). The full list of options can be found here.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// It is required if type is Webhook. This is the url Opsgenie will be sending request to.
	WebhookURL *string `json:"webhookUrl,omitempty" tf:"webhook_url,omitempty"`
}

type IntegrationObservation struct {
	AllowConfigurationAccess *bool `json:"allowConfigurationAccess,omitempty" tf:"allow_configuration_access,omitempty"`

	// This parameter is for configuring the write access of integration. If write access is restricted, the integration will not be authorized to write within any domain. Default: true.
	AllowWriteAccess *bool `json:"allowWriteAccess,omitempty" tf:"allow_write_access,omitempty"`

	// This parameter is for specifying whether the integration will be enabled or not. Default: true
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// The id of the responder.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If enabled, the integration will ignore recipients sent in request payloads. Default: false.
	IgnoreRespondersFromPayload *bool `json:"ignoreRespondersFromPayload,omitempty" tf:"ignore_responders_from_payload,omitempty"`

	// Name of the integration. Name must be unique for each integration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner team id of the integration. If changed, this will recreate a new API integration, which will probably have a different API key.
	OwnerTeamID *string `json:"ownerTeamId,omitempty" tf:"owner_team_id,omitempty"`

	// User, schedule, teams or escalation names to calculate which users will receive the notifications of the alert.
	Responders []RespondersObservation `json:"responders,omitempty" tf:"responders,omitempty"`

	// If enabled, notifications that come from alerts will be suppressed. Default: false.
	SuppressNotifications *bool `json:"suppressNotifications,omitempty" tf:"suppress_notifications,omitempty"`

	// Type of the integration (API, Marid, Prometheus, etc). The full list of options can be found here.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// It is required if type is Webhook. This is the url Opsgenie will be sending request to.
	WebhookURL *string `json:"webhookUrl,omitempty" tf:"webhook_url,omitempty"`
}

type IntegrationParameters struct {

	// +kubebuilder:validation:Optional
	AllowConfigurationAccess *bool `json:"allowConfigurationAccess,omitempty" tf:"allow_configuration_access,omitempty"`

	// This parameter is for configuring the write access of integration. If write access is restricted, the integration will not be authorized to write within any domain. Default: true.
	// +kubebuilder:validation:Optional
	AllowWriteAccess *bool `json:"allowWriteAccess,omitempty" tf:"allow_write_access,omitempty"`

	// This parameter is for specifying whether the integration will be enabled or not. Default: true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// If enabled, the integration will ignore recipients sent in request payloads. Default: false.
	// +kubebuilder:validation:Optional
	IgnoreRespondersFromPayload *bool `json:"ignoreRespondersFromPayload,omitempty" tf:"ignore_responders_from_payload,omitempty"`

	// Name of the integration. Name must be unique for each integration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner team id of the integration. If changed, this will recreate a new API integration, which will probably have a different API key.
	// +kubebuilder:validation:Optional
	OwnerTeamID *string `json:"ownerTeamId,omitempty" tf:"owner_team_id,omitempty"`

	// User, schedule, teams or escalation names to calculate which users will receive the notifications of the alert.
	// +kubebuilder:validation:Optional
	Responders []RespondersParameters `json:"responders,omitempty" tf:"responders,omitempty"`

	// If enabled, notifications that come from alerts will be suppressed. Default: false.
	// +kubebuilder:validation:Optional
	SuppressNotifications *bool `json:"suppressNotifications,omitempty" tf:"suppress_notifications,omitempty"`

	// Type of the integration (API, Marid, Prometheus, etc). The full list of options can be found here.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// It is required if type is Webhook. This is the url Opsgenie will be sending request to.
	// +kubebuilder:validation:Optional
	WebhookURL *string `json:"webhookUrl,omitempty" tf:"webhook_url,omitempty"`
}

type RespondersInitParameters struct {

	// The id of the responder.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of the integration (API, Marid, Prometheus, etc). The full list of options can be found here.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RespondersObservation struct {

	// The id of the responder.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of the integration (API, Marid, Prometheus, etc). The full list of options can be found here.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RespondersParameters struct {

	// The id of the responder.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of the integration (API, Marid, Prometheus, etc). The full list of options can be found here.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationParameters `json:"forProvider"`
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
	InitProvider IntegrationInitParameters `json:"initProvider,omitempty"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Integration is the Schema for the Integrations API. Manages an API Integration within Opsgenie.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IntegrationSpec   `json:"spec"`
	Status IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationList contains a list of Integrations
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

// Repository type metadata.
var (
	Integration_Kind             = "Integration"
	Integration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Integration_Kind}.String()
	Integration_KindAPIVersion   = Integration_Kind + "." + CRDGroupVersion.String()
	Integration_GroupVersionKind = CRDGroupVersion.WithKind(Integration_Kind)
)

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
