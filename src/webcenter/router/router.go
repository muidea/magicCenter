package router

import (
	"martini"
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

func init() {
	martiniRouter = martini.NewRouter()
}

func Router() martini.Router {
	return martiniRouter
}

func AddGetRoute(pattern string, handler interface{}) {
	martiniRouter.Get(pattern, handler)	
}

func RemoveGetRoute(pattern string) {
	
}

func AddPostRoute(pattern string, handler interface{}) {
	martiniRouter.Post(pattern,handler)
}

func RemovePostRoute(pattern string) {
	
}


