package postgresflex

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_postgresflex_instance", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("stackit_postgresflex_user", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_postgresflex_instance",
		}
	})
}
