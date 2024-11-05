package alertpolicy

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_alert_policy", func(r *config.Resource) {
		r.Kind = "AlertPolicy"

		r.References["team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})
}
