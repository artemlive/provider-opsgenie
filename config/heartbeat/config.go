package heartbeat

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_heartbeat", func(r *config.Resource) {
		r.Kind = "Heartbeat"

		r.References["owner_team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})
}
