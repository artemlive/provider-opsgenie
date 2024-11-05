package maintenance

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_maintenance", func(r *config.Resource) {
		r.Kind = "Maintenance"

		r.References["rules.entity.id"] = config.Reference{
			TerraformName: "opsgenie_email_integration",
		}
	})
}
