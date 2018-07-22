package route

import (
	"muidea.com/magicCenter/common"
)

// AppendEndpointRoute append endpoint route
func AppendEndpointRoute(routes []common.Route, endpointHandler common.EndpointHandler, accountHandler common.AccountHandler, moduleHub common.ModuleHub, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateQueryEndpointRoute(endpointHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateQueryByIDEndpointRoute(endpointHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreatePostEndpointRoute(endpointHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateDeleteEndpointRoute(endpointHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreatePutEndpointRoute(endpointHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetEndpointAuthRoute(endpointHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}
