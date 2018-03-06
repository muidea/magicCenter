package route

import (
	"muidea.com/magicCenter/application/common"
)

// AppendAuthorityRoute append authority route
func AppendAuthorityRoute(routes []common.Route, authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) []common.Route {

	rt := CreateGetACLByIDRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateGetACLByModuleRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreatePostACLRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateDeleteACLRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreatePutACLRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateGetACLAuthGroupRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreatePutACLAuthGroupRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateGetModuleACLRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateGetModuleUserAuthGroupRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreatePutModuleUserAuthGroupRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetUserModuleAuthGroupRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreatePutUserModuleAuthGroupRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetUserACLRoute(authorityHandler, accountHandler)
	routes = append(routes, rt)

	return routes
}
