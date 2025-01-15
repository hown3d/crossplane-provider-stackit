package postgresflex

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/stackitcloud/crossplane-provider-stackit/config/postgresflex"
)

func atProviderExtractorPath(attribute string) string {
	return fmt.Sprintf("%s.AtProviderExtractor(\"%s\")", SelfPackagePath, attribute)
}

func AtProviderExtractor(attribute string) reference.ExtractValueFn {
	return func(m xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(m)
		if err != nil {
			return ""
		}
		r, err := paved.GetString(fmt.Sprintf("status.atProvider.%s", attribute))
		if err != nil {
			return ""
		}
		return r
	}
}

func UserNameExtractor() reference.ExtractValueFn {
	return func(m xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(m)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.username")
		if err != nil {
			return ""
		}
		return r
	}
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_postgresflex_instance", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("stackit_postgresflex_user", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_postgresflex_instance",
			Extractor:     atProviderExtractorPath("instanceId"),
		}
	})
	p.AddResourceConfigurator("stackit_postgresflex_database", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_postgresflex_instance",
			Extractor:     atProviderExtractorPath("instanceId"),
		}
		r.References["owner"] = config.Reference{
			TerraformName: "stackit_postgresflex_user",
			Extractor:     atProviderExtractorPath("username"),
		}
	})
}
