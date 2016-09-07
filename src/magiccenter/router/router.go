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

// NewRoute 新建一个路由对象
func NewRoute(rType, rPattern string, rHandler interface{}, rVerifier interface{}) common.Route {
	r := route{}
	r.rType = rType
	r.rPattern = rPattern
	r.rHandler = rHandler
	r.rVerifier = rVerifier

	return &r
}

var martiniRouter martini.Router
var routerVerifier map[string]interface{}

func init() {
	martiniRouter = martini.NewRouter()

	routerVerifier = make(map[string]interface{})
}

// Router 返回系统Router
func Router() martini.Router {
	return martiniRouter
}

// AddGetRoute 添加一条Get路由
func AddGetRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("pattern[get]:%s", pattern)
	}

	martiniRouter.Get(pattern, handler)
}

// RemoveGetRoute 清除一条Get路由
func RemoveGetRoute(pattern string) {
	delete(routerVerifier, pattern)
}

// AddPutRoute 添加一条Put路由
func AddPutRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("pattern[put]:%s", pattern)
	}

	martiniRouter.Put(pattern, handler)
}

// RemovePutRoute 清除一条Put路由
func RemovePutRoute(pattern string) {
	delete(routerVerifier, pattern)
}

// AddPostRoute 添加一条Post路由
func AddPostRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("pattern[post]:%s", pattern)
	}

	martiniRouter.Post(pattern, handler)
}

// RemovePostRoute 清除一条Post路由
func RemovePostRoute(pattern string) {
	delete(routerVerifier, pattern)
}

// AddDeleteRoute 添加一条Delete路由
func AddDeleteRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		routerVerifier[pattern] = verifier
	}

	if martini.Env != martini.Prod {
		log.Printf("pattern[delete]:%s", pattern)
	}

	martiniRouter.Delete(pattern, handler)
}

// RemoveDeleteRoute 清除一条Delete路由
func RemoveDeleteRoute(pattern string) {
	delete(routerVerifier, pattern)
}

// AddAuthVerifier 添加一个路由的权限校验器
func AddAuthVerifier(pattern string, verifier interface{}) {
	util.ValidateFunc(verifier)
	routerVerifier[pattern] = verifier
}

// RemoveAuthVerifier 清除一个路由的权限校验器
func RemoveAuthVerifier(pattern string) {
	delete(routerVerifier, pattern)
}

// VerifyAuthority 校验路由权限
func VerifyAuthority(res http.ResponseWriter, req *http.Request) bool {
	verifiter, found := routerVerifier[req.URL.Path]
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
