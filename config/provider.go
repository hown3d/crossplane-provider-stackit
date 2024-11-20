/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/stackitcloud/crossplane-provider-stackit/config/postgresflex"
)

const (
	resourcePrefix = "stackit"
	modulePath     = "github.com/stackitcloud/crossplane-provider-stackit"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stackit.cloud"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions())

	for _, configure := range []func(provider *ujconfig.Provider){
		postgresflex.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
