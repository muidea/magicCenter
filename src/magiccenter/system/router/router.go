package router

import (
	"log"
	"magiccenter/common"
	"martini"
	"net/http"
	"reflect"

	"muidea.com/util"
)

type route struct {
	rType     string
	rPattern  string
	rHandler  interface{}
	rVerifier interface{}
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

func (r *route) Verifier() interface{} {
	return r.rVerifier
}

// CreateRouter 新建Router
func CreateRouter() common.Router {
	impl := impl{}
	impl.martiniRouter = martini.NewRouter()
	impl.routerVerifier = make(map[string]interface{})

	return &impl
}

// NewRoute 新建一个路由对象
func (instance impl) NewRoute(rType, rPattern string, rHandler interface{}, rVerifier interface{}) common.Route {
	r := route{}
	r.rType = rType
	r.rPattern = rPattern
	r.rHandler = rHandler
	r.rVerifier = rVerifier

	return &r
}

// AddRoute 增加Route
func (instance impl) AddRoute(rt common.Route) {
	switch rt.Type() {
	case common.GET:
		instance.AddGetRoute(rt.Pattern(), rt.Handler(), rt.Verifier())
	case common.POST:
		instance.AddPostRoute(rt.Pattern(), rt.Handler(), rt.Verifier())
	case common.DELETE:
		instance.AddDeleteRoute(rt.Pattern(), rt.Handler(), rt.Verifier())
	case common.PUT:
		instance.AddPutRoute(rt.Pattern(), rt.Handler(), rt.Verifier())
	}
}

// RemoveRoute 清除Route
func (instance impl) RemoveRoute(rt common.Route) {
	switch rt.Type() {
	case common.GET:
		instance.RemoveGetRoute(rt.Pattern())
	case common.POST:
		instance.RemovePostRoute(rt.Pattern())
	case common.DELETE:
		instance.RemoveDeleteRoute(rt.Pattern())
	case common.PUT:
		instance.RemovePutRoute(rt.Pattern())
	}
}

// Router 返回系统Router
func (instance impl) Router() martini.Router {
	return instance.martiniRouter
}

// AddGetRoute 添加一条Get路由
func (instance impl) AddGetRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		instance.routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("[get]:%s", pattern)
	}

	instance.martiniRouter.Get(pattern, handler)
}

// RemoveGetRoute 清除一条Get路由
func (instance impl) RemoveGetRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// AddPutRoute 添加一条Put路由
func (instance impl) AddPutRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		instance.routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("[put]:%s", pattern)
	}

	instance.martiniRouter.Put(pattern, handler)
}

// RemovePutRoute 清除一条Put路由
func (instance impl) RemovePutRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// AddPostRoute 添加一条Post路由
func (instance impl) AddPostRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		instance.routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("[post]:%s", pattern)
	}

	instance.martiniRouter.Post(pattern, handler)
}

// RemovePostRoute 清除一条Post路由
func (instance impl) RemovePostRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// AddDeleteRoute 添加一条Delete路由
func (instance impl) AddDeleteRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		instance.routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("[delete]:%s", pattern)
	}

	instance.martiniRouter.Delete(pattern, handler)
}

// RemoveDeleteRoute 清除一条Delete路由
func (instance impl) RemoveDeleteRoute(pattern string) {
	delete(instance.routerVerifier, pattern)
}

// Dispatch 分发一次请求
func (instance impl) Dispatch(res http.ResponseWriter, req *http.Request) {

}

// VerifyAuthority 校验路由权限
func (instance impl) VerifyAuthority(res http.ResponseWriter, req *http.Request) bool {
	verifiter, found := instance.routerVerifier[req.URL.Path]
	if !found {
		// 找不到verifier,说明不需要进权限校验，返回校验通过
		return true
	}

	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(res)
	in[1] = reflect.ValueOf(req)
	value := reflect.ValueOf(verifiter).Call(in)[0]
	return value.Bool()
}

// AddAuthVerifier 添加一个路由的权限校验器
func (instance impl) AddAuthVerifier(pattern string, verifier interface{}) {
	util.ValidateFunc(verifier)
	instance.routerVerifier[pattern] = verifier
}

// RemoveAuthVerifier 清除一个路由的权限校验器
func (instance impl) RemoveAuthVerifier(pattern string) {
	delete(instance.routerVerifier, pattern)
}
