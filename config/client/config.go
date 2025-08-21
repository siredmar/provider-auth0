package client

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_client", func(r *config.Resource) {
		r.ShortGroup = "client"
	})
}
