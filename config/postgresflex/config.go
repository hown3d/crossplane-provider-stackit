package postgresflex

import (
	"fmt"
	"math"
	"strconv"

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
		r.Sensitive.AdditionalConnectionDetailsFn = additionalConnectionDetails("host", "port", "password", "username")
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_postgresflex_instance",
			Extractor:     atProviderExtractorPath("instanceId"),
		}
	})
	p.AddResourceConfigurator("stackit_postgresflex_database", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = additionalConnectionDetails("name")

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

func additionalConnectionDetails(keys ...string) config.AdditionalConnectionDetailsFn {
	return func(attr map[string]any) (map[string][]byte, error) {
		conn := map[string][]byte{}
		for _, k := range keys {
			val := attr[k]
			switch typedVal := val.(type) {
			case string:
				conn[k] = []byte(typedVal)
			case int:
				conn[k] = []byte(strconv.Itoa(typedVal))
			case float64:
				if _, frac := math.Modf(typedVal); frac == 0 {
					conn[k] = []byte(strconv.Itoa(int(typedVal)))
				} else {
					conn[k] = []byte(strconv.FormatFloat(typedVal, 'f', -1, 64))
				}
			default:
				return nil, fmt.Errorf("unknown type for key %s: %T", k, val)
			}
		}
		return conn, nil
	}
}
