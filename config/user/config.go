package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_user", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "User"
	})
	p.AddResourceConfigurator("opsgenie_user_contact", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "UserContact"

		r.References["username"] = config.Reference{
			TerraformName: "opsgenie_user",
		}
	})
}
