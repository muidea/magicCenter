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

	rt, _ = CreateAccountChangePasswordRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateEndpointLoginRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateEndpointLogoutRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateEndpointStatusRoute(casHandler, sessionRegistry)
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

// CreateAccountChangePasswordRoute 创建AccountChangePassword Route
func CreateAccountChangePasswordRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := accountChangePasswordRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateEndpointLoginRoute 创建EndpointVerify Route
func CreateEndpointLoginRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := endpointLoginRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateEndpointLogoutRoute 创建EndpointLogout Route
func CreateEndpointLogoutRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := endpointLogoutRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}

// CreateEndpointStatusRoute 创建EndpointStatus Route
func CreateEndpointStatusRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	i := endpointStatusRoute{
		casHandler:      casHandler,
		sessionRegistry: sessionRegistry}
	return &i, true
}
