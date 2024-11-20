package postgresflex

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/stackitcloud/crossplane-provider-stackit/config/postgresflex"

	// PathInstanceIDExtractor is the golang path to InstanceIDExtractor function
	// in this package.
	PathInstanceIDExtractor = SelfPackagePath + ".InstanceIDExtractor()"
)

func InstanceIDExtractor() reference.ExtractValueFn {
	return func(m xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(m)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.instanceId")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_postgresflex_instance", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("stackit_postgresflex_user", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_postgresflex_instance",
			Extractor:     PathInstanceIDExtractor,
		}
	})
}
