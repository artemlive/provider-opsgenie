package escalation

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_escalation", func(r *config.Resource) {
		r.Kind = "Escalation"

		// TODO: https://github.com/crossplane/upjet/issues/414
		// it's impossible to have a multiple types for the same field reference
		// I want to try to achieve the same behavior through the checking the type field
		// e.g. if type == team, we should use team reference, etc
		r.References["rules.recipient.id"] = config.Reference{
			TerraformName: "opsgenie_schedule",
		}
	})
}
