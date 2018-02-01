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

	rt = CreateModuleACLGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateModuleUserGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateModuleUserPutRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateUserModuleGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateUserModulePutRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateUserAuthGroupGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateUserAuthGroupPutRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateUserACLGetRoute(authorityHandler)
	routes = append(routes, rt)

	return routes
}
