package service

/*
路由器

管理系统中的路由信息

1、新建一条路由规则
2、增加删除路由信息
3、分发请求
*/

import (
	"net/http"

	"muidea.com/magicCenter/application/common"

	"github.com/go-martini/martini"
)

// Router 路由器对象
type Router interface {

	// 新建Route
	NewRoute(rType, rPattern string, rHandler interface{}, rVerifier interface{}) common.Route
	// 增加路由
	AddRoute(baseURL string, rt common.Route)
	// 清除路由
	RemoveRoute(baseURL string, rt common.Route)

	// 增加Get路由
	AddGetRoute(pattern string, handler, verifier interface{})
	// 清除Get路由
	RemoveGetRoute(pattern string)
	// 增加Post路由
	AddPostRoute(pattern string, handler, verifier interface{})
	// 清除Post路由
	RemovePostRoute(pattern string)
	// 增加Delete路由
	AddDeleteRoute(pattern string, handler, verifier interface{})
	// 清除Delete路由
	RemoveDeleteRoute(pattern string)
	// 增加Put路由
	AddPutRoute(pattern string, handler, verifier interface{})
	// 清除Put路由
	RemovePutRoute(pattern string)

	// 返回Martini.Router对象
	Router() martini.Router

	// 分发一条请求
	Dispatch(res http.ResponseWriter, req *http.Request)
	// 校验权限，无权限返回false，有权限返回true
	VerifyAuthority(res http.ResponseWriter, req *http.Request) bool
}
