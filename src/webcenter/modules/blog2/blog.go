package blog2

import (
	"log"
	"webcenter/module"
)

const ID = "f17133ec-63e9-4b46-8757-e6ca1af6fe4e"
const URI = "/blog2"

type blog2 struct {
	
}

var instance *blog2 = nil

func init() {
	log.Println("register blog2 module")
	
	instance = &blog2{}
	
	module.RegisterModule(instance)
}

func (this *blog2) Startup() {
}

func (this *blog2) Cleanup() {
	
}

func (this *blog2) ID() string {
	return ID
}

func (this *blog2) Uri() string {
	return URI
}

func (this *blog2) Routes() []module.Route {
	routes := []module.Route{module.NewRoute(module.GET,"/view/",viewArticleHandler), module.NewRoute(module.GET,"/",indexHandler)}
	
	return routes
}


