package route

import (
	"muidea.com/magicCenter/application/common"
)

// AppendAuthorityRoute append authority route
func AppendAuthorityRoute(routes []common.Route, authorityHandler common.AuthorityHandler) []common.Route {

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

	rt = CreateGetModuleUserAuthGroupRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreatePutModuleUserAuthGroupRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateGetUserModuleAuthGroupRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreatePutUserModuleAuthGroupRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateGetUserACLRoute(authorityHandler)
	routes = append(routes, rt)

	return routes
}
