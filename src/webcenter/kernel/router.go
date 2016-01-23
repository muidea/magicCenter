package kernel

import (
	"martini"
)

type router struct {
	martiniRouter martini.Router
}

func NewRouter() *router {
	r := router{}
	
	r.martiniRouter = martini.NewRouter()
	
	return &r
}

func (this *router) Router() martini.Router {
	return this.martiniRouter
}

func (this *router)AddGetRoute(pattern string, handler interface{}) {
	this.martiniRouter.Get(pattern, handler)
}

func (this *router)RemoveGetRoute(pattern string) {
	
}

func (this *router)AddPostRoute(pattern string, handler interface{}) {
	this.martiniRouter.Post(pattern,handler)
}

func (this *router)RemovePostRoute(pattern string) {
	
}

