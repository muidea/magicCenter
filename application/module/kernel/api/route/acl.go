package route

import "muidea.com/magicCenter/application/common"

// AppendACLRoute 追加Acl Route
func AppendACLRoute(routes []common.Route, modHub common.ModuleHub, sessionRegistry common.SessionRegistry) []common.Route {
	rt, _ := CreateGetACLRoute(modHub)
	routes = append(routes, rt)

	return routes
}

// CreateGetACLRoute GetAcl Route
func CreateGetACLRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := articleGetByIDRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}
