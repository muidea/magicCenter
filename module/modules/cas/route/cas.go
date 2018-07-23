package route

import (
	"muidea.com/magicCenter/common"
)

// AppendAccountRoute 追加account 路由
func AppendAccountRoute(routes []common.Route, casHandler common.CASHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt, _ := CreateAccountLoginRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateAccountLogoutRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateAccountStatusRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateEndpointVerifyRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateAccountLoginRoute 创建AccountLogin Route
func CreateAccountLoginRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := accountLoginRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}

	return &i, true
}

// CreateAccountLogoutRoute 创建AccountLogout Route
func CreateAccountLogoutRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := accountLogoutRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateAccountStatusRoute 创建AccountStatus Route
func CreateAccountStatusRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := accountStatusRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateEndpointVerifyRoute 创建EndpointVerify Route
func CreateEndpointVerifyRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := endpointVerifyRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}
