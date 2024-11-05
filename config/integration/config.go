package integration

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_integration_action", func(r *config.Resource) {
		r.Kind = "IntegrationAction"

		r.References["integration_id"] = config.Reference{
			TerraformName: "opsgenie_api_integration",
		}
		// TODO: https://github.com/crossplane/upjet/issues/414
		// it's impossible to have a multiple types for the same field reference
		// so we support only team reference for now
		// should support: team, user, escalation
		// I want to try to achieve the same behavior through the checking the type field
		// e.g. if type == team, we should use team reference, etc
		r.References["create.responders.id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})

	// TODO: https://github.com/crossplane/upjet/issues/414
	// it's impossible to have a multiple types for the same field reference
	// so we support only team reference for now
	p.AddResourceConfigurator("opsgenie_email_integration", func(r *config.Resource) {
		r.Kind = "EmailIntegration"
		r.References["responders.id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
		r.References["owner_team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})

	// TODO: https://github.com/crossplane/upjet/issues/414
	// it's impossible to have a multiple types for the same field reference
	p.AddResourceConfigurator("opsgeine_api_integration", func(r *config.Resource) {
		r.Kind = "APIIntegration"
		r.References["responders.id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
		r.References["owner_team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})
}
