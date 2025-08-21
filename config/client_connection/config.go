package client_connection

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_client_connection", func(r *config.Resource) {
		r.ShortGroup = "client_connection"
	})
}
