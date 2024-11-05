package service

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_service", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "Service"

		r.References["team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})
	p.AddResourceConfigurator("opsgenie_service_incident_rule", func(r *config.Resource) {
		r.ShortGroup = "service"
		r.Kind = "ServiceIncidentRule"

		r.References["service_id"] = config.Reference{
			TerraformName: "opsgenie_service",
		}

	})
}
