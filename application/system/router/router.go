package router

import (
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/service"
	"muidea.com/magicCenter/foundation/util"

	"github.com/go-martini/martini"
)

type route struct {
	rType    string
	rPattern string
	rHandler interface{}
}

// impl 路由器实现
type impl struct {
	martiniRouter  martini.Router
	routerVerifier map[string]interface{}
}

// Type 路由行为GET/POST
func (r *route) Type() string {
	return r.rType
}

// Pattern 路由规则
func (r *route) Pattern() string {
	return r.rPattern
}

// Handler 路由处理器
func (r *route) Handler() interface{} {
	return r.rHandler
}

// CreateRouter 新建Router
func CreateRouter() service.Router {
	impl := impl{}
	impl.martiniRouter = martini.NewRouter()

	return &impl
}

// NewRoute 新建一个路由对象
func (instance *impl) NewRoute(rType, rPattern string, rHandler interface{}) common.Route {
	r := route{}
	r.rType = rType
	r.rPattern = rPattern
	r.rHandler = rHandler

	return &r
}

// AddRoute 增加Route
func (instance *impl) AddRoute(baseURL string, rt common.Route) {
	fullURL := util.JoinURL(baseURL, rt.Pattern())
	switch rt.Type() {
	case common.GET:
		instance.AddGetRoute(fullURL, rt.Handler())
	case common.POST:
		instance.AddPostRoute(fullURL, rt.Handler())
	case common.DELETE:
		instance.AddDeleteRoute(fullURL, rt.Handler())
	case common.PUT:
		instance.AddPutRoute(fullURL, rt.Handler())
	}
}

// RemoveRoute 清除Route
func (instance *impl) RemoveRoute(baseURL string, rt common.Route) {
	fullURL := util.JoinURL(baseURL, rt.Pattern())
	switch rt.Type() {
	case common.GET:
		instance.RemoveGetRoute(fullURL)
	case common.POST:
		instance.RemovePostRoute(fullURL)
	case common.DELETE:
		instance.RemoveDeleteRoute(fullURL)
	case common.PUT:
		instance.RemovePutRoute(fullURL)
	}
}

// Router 返回系统Router
func (instance *impl) Router() martini.Router {
	return instance.martiniRouter
}

// AddGetRoute 添加一条Get路由
func (instance *impl) AddGetRoute(pattern string, handler interface{}) {
	if martini.Env != martini.Prod {
		log.Printf("[get]:%s", pattern)
	}

	instance.martiniRouter.Get(pattern, handler)
}

// RemoveGetRoute 清除一条Get路由
func (instance *impl) RemoveGetRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// AddPutRoute 添加一条Put路由
func (instance *impl) AddPutRoute(pattern string, handler interface{}) {
	if martini.Env != martini.Prod {
		log.Printf("[put]:%s", pattern)
	}

	instance.martiniRouter.Put(pattern, handler)
}

// RemovePutRoute 清除一条Put路由
func (instance *impl) RemovePutRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// AddPostRoute 添加一条Post路由
func (instance *impl) AddPostRoute(pattern string, handler interface{}) {
	if martini.Env != martini.Prod {
		log.Printf("[post]:%s", pattern)
	}

	instance.martiniRouter.Post(pattern, handler)
}

// RemovePostRoute 清除一条Post路由
func (instance *impl) RemovePostRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// AddDeleteRoute 添加一条Delete路由
func (instance *impl) AddDeleteRoute(pattern string, handler interface{}) {
	if martini.Env != martini.Prod {
		log.Printf("[delete]:%s", pattern)
	}

	instance.martiniRouter.Delete(pattern, handler)
}

// RemoveDeleteRoute 清除一条Delete路由
func (instance *impl) RemoveDeleteRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// Dispatch 分发一次请求
func (instance *impl) Dispatch(res http.ResponseWriter, req *http.Request) {

}
