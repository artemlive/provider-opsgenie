package incident

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_incident_template", func(r *config.Resource) {
		r.Kind = "IncidentTemplate"

		r.References["impacted_services"] = config.Reference{
			TerraformName: "opsgenie_service",
		}
	})
}
