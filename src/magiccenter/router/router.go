package router

import (
	"net/http"
	"reflect"
	"martini"
	"muidea.com/util"
)

type Route interface {
	Type() string
	Pattern() string
	Handler() interface{}
}

type route struct {
	rType string
	rPattern string
	rHandler interface{}
}

func (this *route) Type() string {
	return this.rType
}

func (this *route) Pattern() string {
	return this.rPattern
}

func (this *route) Handler() interface{} {
	return this.rHandler
}

func NewRoute(rType, rPattern string, rHandler interface{}) Route {
	r := route{}
	r.rType = rType
	r.rPattern = rPattern
	r.rHandler = rHandler
	
	return &r	
}

var martiniRouter martini.Router
var routerVerifier map[string]interface{}

func init() {
	martiniRouter = martini.NewRouter()
	
	routerVerifier = make(map[string]interface{})
}

func Router() martini.Router {
	return martiniRouter
}

func AddGetRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		routerVerifier[pattern] = verifier
	}
	
	martiniRouter.Get(pattern, handler)	
}

func RemoveGetRoute(pattern string) {
	delete(routerVerifier, pattern)
}

func AddPostRoute(pattern string, handler, verifier interface{}) {
	// 如果verifier为nil则表示不需要进行权限校验
	// 所以在verifier为nil的情况下不需要校验verifier是否为func
	if verifier != nil {
		util.ValidateFunc(verifier)
		routerVerifier[pattern] = verifier
	}
		
	martiniRouter.Post(pattern,handler)
}

func RemovePostRoute(pattern string) {
	delete(routerVerifier, pattern)
}

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



