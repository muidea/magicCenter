package common

import (
	"muidea.com/magicCenter/application/common/model"
)

// 基本HTTP行为定义
const (
	GET    = "GET"
	PUT    = "PUT"
	POST   = "POST"
	DELETE = "DELETE"
)

// Route 路由接口
type Route interface {
	// Action 路由行为GET/PUT/POST/DELETE
	Method() string
	// Pattern 路由规则, 以'/'开始
	Pattern() string
	// Handler 路由处理器
	Handler() interface{}
}

type route struct {
	rRoute   model.Route
	rHandler interface{}
}

// Type 路由行为GET/POST
func (r *route) Method() string {
	return r.rRoute.Method
}

// Pattern 路由规则
func (r *route) Pattern() string {
	return r.rRoute.Pattern
}

// Handler 路由处理器
func (r *route) Handler() interface{} {
	return r.rHandler
}

// NewRoute 新建一个路由对象
func NewRoute(method, pattern string, rHandler interface{}) Route {
	r := route{rRoute: model.Route{Method: method, Pattern: pattern}, rHandler: rHandler}

	return &r
}
