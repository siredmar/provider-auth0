package client_connection_grant

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_client_connection_grant", func(r *config.Resource) {
		r.ShortGroup = "client_connection_grant"
	})
}
