package blog

import (
	"log"
	"webcenter/module"	
)

const ID = "f17133ec-63e9-4b46-8757-e6ca1af6fe3e"
const URI = "/blog"

type blog struct {
	
}

var instance *blog = nil

func init() {
	log.Println("register blog module")
	
	instance = &blog{}
	
	module.RegisterModule(instance)
}

func (this *blog) Startup() {
}

func (this *blog) Cleanup() {
	
}

func (this *blog) ID() string {
	return ID
}

func (this *blog) Uri() string {
	return URI
}

func (this *blog) Routes() []module.Route {
	routes := []module.Route{module.NewRoute(module.GET,"/view/",viewArticleHandler), module.NewRoute(module.GET,"/",indexHandler)}
	
	return routes
}


