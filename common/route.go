package common

import (
	"muidea.com/magicCommon/model"
)

// 基本HTTP行为定义
const (
	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
)

// Route 路由接口
type Route interface {
	// Action 路由行为GET/PUT/POST/DELETE
	Method() string
	// Pattern 路由规则, 以'/'开始
	Pattern() string
	// Handler 路由处理器
	Handler() interface{}
	// AuthGroup 路由所需的授权组
	AuthGroup() int
}

type route struct {
	rRoute     model.Route
	rHandler   interface{}
	rAuthGroup int
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

func (r *route) AuthGroup() int {
	return r.rAuthGroup
}

// NewRoute 新建一个路由对象
func NewRoute(method, pattern string, rHandler interface{}, authGroup int) Route {
	r := route{rRoute: model.Route{Method: method, Pattern: pattern}, rHandler: rHandler, rAuthGroup: authGroup}

	return &r
}
