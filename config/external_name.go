/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"opsgenie_user":                  config.IdentifierFromProvider,
	"opsgenie_user_contact":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.username }}/{{ .external_name }}"),
	"opsgenie_team":                  config.IdentifierFromProvider,
	"opsgenie_team_routing_rule":     config.TemplatedStringAsIdentifier("name", "{{ .parameters.team }}/{{ .external_name }}"),
	"opsgenie_service":               config.IdentifierFromProvider,
	"opsgenie_service_incident_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.service }}/{{ .external_name }}"),
	"opsgenie_schedule":              config.IdentifierFromProvider,
	"opsgenie_schedule_rotation":     config.TemplatedStringAsIdentifier("name", "{{ .parameters.schedule }}/{{ .external_name }}"),
	"opsgenie_notification_rule":     config.TemplatedStringAsIdentifier("name", "{{ .parameters.user }}/{{ .external_name }}"),
	"opsgenie_notification_policy":   config.TemplatedStringAsIdentifier("name", "{{ .parameters.team }}/{{ .external_name }}"),
	"opsgenie_maintenance":           config.IdentifierFromProvider,
	"opsgenie_integration_action":    config.IdentifierFromProvider,
	"opsgenie_incident_opsgenie":     config.IdentifierFromProvider,
	"opsgenie_heartbeat":             config.IdentifierFromProvider,
	"opsgenie_escalation":            config.IdentifierFromProvider,
	"opsgenie_email_integration":     config.IdentifierFromProvider,
	"opsgenie_custom_role":           config.IdentifierFromProvider,
	"opsgenie_api_integration":       config.IdentifierFromProvider,
	"opsgenie_alert_policy":          config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
