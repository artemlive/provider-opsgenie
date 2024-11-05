package notification

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_notification_rule", func(r *config.Resource) {
		r.ShortGroup = "notification"
		r.Kind = "NotificationRule"

		r.References["username"] = config.Reference{
			TerraformName: "opsgenie_user",
		}
	})
	p.AddResourceConfigurator("opsgenie_notification_policy", func(r *config.Resource) {
		r.ShortGroup = "notification"
		r.Kind = "NotificationPolicy"

		r.References["team_id"] = config.Reference{
			TerraformName: "opsgenie_team",
		}
	})
}
