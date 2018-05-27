package route

import (
	"muidea.com/magicCenter/application/common"
)

// AppendAuthorityRoute append authority route
func AppendAuthorityRoute(routes []common.Route, authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler, moduleHub common.ModuleHub, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateQueryACLRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateGetACLByIDRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreatePostACLRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateDeleteACLRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreatePutACLRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreatePutACLsRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateGetACLAuthGroupRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreatePutACLAuthGroupRoute(authorityHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateQueryModuleRoute(authorityHandler, accountHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateGetModuleByIDRoute(authorityHandler, accountHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreatePutModuleRoute(authorityHandler, accountHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateQueryUserRoute(authorityHandler, accountHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateGetUserByIDRoute(authorityHandler, accountHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreatePutUserRoute(authorityHandler, accountHandler, moduleHub)
	routes = append(routes, rt)

	rt = CreateQueryEndpointRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateQueryByIDEndpointRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreatePostEndpointRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateDeleteEndpointRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreatePutEndpointRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetEndpointAuthRoute(authorityHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}
