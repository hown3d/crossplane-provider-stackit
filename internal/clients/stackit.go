/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/stackitcloud/crossplane-provider-stackit/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal stackit credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		tfSetup := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return tfSetup, errors.New(errNoProviderConfig)
		}
		providerConfig := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, providerConfig); err != nil {
			return tfSetup, errors.Wrap(err, errGetProviderConfig)
		}

		tracker := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := tracker.Track(ctx, mg); err != nil {
			return tfSetup, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, providerConfig.Spec.Credentials.Source, client, providerConfig.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return tfSetup, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return tfSetup, errors.Wrap(err, errUnmarshalCredentials)
		}

		tfSetup.Configuration = map[string]any{
			"service_account_email": creds["service_account_email"],
			"service_account_token": creds["service_account_token"],
			"region":                providerConfig.Spec.Region,
		}

		return tfSetup, nil
	}
}
