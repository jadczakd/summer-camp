// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/jadczakd/summer-camp/summer_app/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.SummerAppService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateEndpoint: MakeCreateEndpoint(s),
		HealthEndpoint: MakeHealthEndpoint(s),
	}
	for _, m := range mdw["Create"] {
		eps.CreateEndpoint = m(eps.CreateEndpoint)
	}
	for _, m := range mdw["Health"] {
		eps.HealthEndpoint = m(eps.HealthEndpoint)
	}
	return eps
}
