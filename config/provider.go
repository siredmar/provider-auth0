/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/siredmar/provider-auth0/config/client"
	"github.com/siredmar/provider-auth0/config/client_connection"
	"github.com/siredmar/provider-auth0/config/client_connection_grant"
)

const (
	resourcePrefix = "auth0"
	modulePath     = "github.com/siredmar/provider-auth0"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		client.Configure,
		client_connection.Configure,
		client_connection_grant.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
