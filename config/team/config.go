package team

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_team", func(r *config.Resource) {
		r.ShortGroup = "team"
		r.Kind = "Team"
		r.References["member.id"] = config.Reference{
			TerraformName: "opsgenie_user",
		}
	})
	p.AddResourceConfigurator("opsgenie_team_routing_rule", func(r *config.Resource) {
		r.ShortGroup = "team"
		r.Kind = "TeamRoutingRule"

		r.References["team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
		r.References["notify.name"] = config.Reference{
			TerraformName: "opsgenie_schedule",
		}
	})
	p.AddResourceConfigurator("opsgenie_custom_role", func(r *config.Resource) {
		r.Kind = "CustomRole"
	})
	// there can be a schedule reference in the team routing rule
	// and a team reference in the schedule, so it generates a cycle
	// that's why we have this resource configurator here
	p.AddResourceConfigurator("opsgenie_schedule", func(r *config.Resource) {
		r.ShortGroup = "team"
		r.Kind = "Schedule"

		r.References["owner_team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})
	p.AddResourceConfigurator("opsgenie_schedule_rotation", func(r *config.Resource) {
		r.ShortGroup = "team"
		r.Kind = "ScheduleRotation"

		r.References["schedule_id"] = config.Reference{
			TerraformName: "opsgenie_schedule",
		}

		r.References["participant.id"] = config.Reference{
			TerraformName: "opsgenie_user",
		}
	})
}
