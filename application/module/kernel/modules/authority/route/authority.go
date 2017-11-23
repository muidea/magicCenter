package route

import (
	"muidea.com/magicCenter/application/common"
)

// AppendAuthorityRoute append authority route
func AppendAuthorityRoute(routes []common.Route, authorityHandler common.AuthorityHandler) []common.Route {

	rt := CreateACLGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLGetByModuleRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLPostRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLDeleteRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLPutRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLAuthGroupGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLAuthGroupPutRoute(authorityHandler)
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
