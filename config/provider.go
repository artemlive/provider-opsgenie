/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/macpaw/provider-opsgenie/config/alertpolicy"
	"github.com/macpaw/provider-opsgenie/config/escalation"
	"github.com/macpaw/provider-opsgenie/config/heartbeat"
	"github.com/macpaw/provider-opsgenie/config/incident"
	"github.com/macpaw/provider-opsgenie/config/integration"
	"github.com/macpaw/provider-opsgenie/config/maintenance"
	"github.com/macpaw/provider-opsgenie/config/notification"
	"github.com/macpaw/provider-opsgenie/config/service"
	"github.com/macpaw/provider-opsgenie/config/team"
	"github.com/macpaw/provider-opsgenie/config/user"
)

const (
	resourcePrefix = "opsgenie"
	modulePath     = "github.com/macpaw/provider-opsgenie"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("macpaw.dev"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		user.Configure,
		team.Configure,
		service.Configure,
		//		schedule.Configure,
		notification.Configure,
		maintenance.Configure,
		integration.Configure,
		incident.Configure,
		heartbeat.Configure,
		escalation.Configure,
		alertpolicy.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
