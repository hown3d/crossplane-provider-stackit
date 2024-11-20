package postgresflex

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

const (
	ErrFmtNoAttribute = "attribute %s not found"
)

func getInstanceId(tfstate map[string]any) (string, error) {
	id, ok := tfstate["instance_id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "instance_id")
	}
	return fmt.Sprintf("%s", id), nil
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_postgresflex_instance", func(r *config.Resource) {
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]any, externalName string) {
				base["name"] = externalName
			},
			GetExternalNameFn: getInstanceId,
			GetIDFn:           config.ExternalNameAsID,
			OmittedFields:     []string{"name"},
		}
	})
	p.AddResourceConfigurator("stackit_postgresflex_user", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "stackit_postgresflex_instance",
		}
	})
}
