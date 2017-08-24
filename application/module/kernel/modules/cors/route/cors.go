package route

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/cors/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendCorsRoute 追加cors 路由
func AppendCorsRoute(routes []common.Route, corsHandler common.CorsHandler) []common.Route {
	rt, _ := CreateCorsCheckRoute(corsHandler)
	routes = append(routes, rt)

	return routes
}

// CreateCorsCheckRoute 创建CorsCheck Route
func CreateCorsCheckRoute(corsHandler common.CorsHandler) (common.Route, bool) {
	i := corsCheckRoute{corsHandler: corsHandler}
	return &i, true
}

type corsCheckRoute struct {
	corsHandler common.CorsHandler
}

func (i *corsCheckRoute) Method() string {
	return common.OPTIONS
}

func (i *corsCheckRoute) Pattern() string {
	return net.JoinURL(def.URL, "**")
}

func (i *corsCheckRoute) Handler() interface{} {
	return i.corsCheckHandler
}

func (i *corsCheckRoute) corsCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
