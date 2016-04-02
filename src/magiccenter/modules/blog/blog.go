package blog

import (
	"log"
	"magiccenter/module"
)

const ID = "f17133ec-63e9-4b46-8757-e6ca1af6fe3e"
const Name = "Blog"
const Description = "blog module"
const URI = "/blog"

type blog struct {
	
}

var instance *blog = nil

func init() {
	log.Println("register blog module")
	
	instance = &blog{}
	
	module.RegisterModule(instance)
}

func (this *blog) ID() string {
	return ID
}

func (this *blog) Name() string {
	return Name
}

func (this *blog) Description() string {
	return Description
}

func (this *blog) Uri() string {
	return URI
}

func (this *blog) Routes() []module.Route {
	//routes := []module.Route{module.NewRoute(module.GET,"/",indexHandler), module.NewRoute(module.GET,"/view/",viewArticleHandler), module.NewRoute(module.GET,"/guestbook/",indexHandler), module.NewRoute(module.GET,"/aboutsite/",indexHandler), module.NewRoute(module.GET,"/aboutme/",indexHandler)}
	routes := []module.Route{}
	
	return routes
}

func (this *blog) Startup() {
}

func (this *blog) Cleanup() {
	
}


