package kernel

import (
	"log"
	"martini"
	"webcenter/module"
	"webcenter/router"
	"webcenter/configuration"
)

var instanceFrame = martini.New()

func RegisterGetHandler(pattern string, h interface{}) {
	router.AddGetRoute(pattern, h)	
}

func UnRegisterGetHandler(pattern string, h interface{}) {
}

func RegisterPostHandler(pattern string, h interface{}) {
	router.AddPostRoute(pattern, h)
}

func UnRegisterPostHandler(pattern string, h interface{}) {
}

func BindStatic(path string) {
	instanceFrame.Use(martini.Static(path))
}

func Initialize() {
	log.Println("initialize kernel...")
	
	staticPath, found := configuration.GetOption(configuration.STATIC_PATH)
	if found {
		BindStatic(staticPath)
	}
	
	resourceFilePath, found := configuration.GetOption(configuration.RESOURCE_PATH)
	if found {
		BindStatic(resourceFilePath)
	}
	
	module.StartupAllModules()
}


func Uninitialize() {
	module.CleanupAllModules()
}

func Run() {
	martiniRouter := router.Router()
	
	instanceFrame.Use(martini.Logger())
	instanceFrame.Use(martini.Recovery())
	instanceFrame.MapTo(martiniRouter, (*martini.Routes)(nil))
	instanceFrame.Action(martiniRouter.Handle)

	martinInstance := &martini.ClassicMartini{instanceFrame, martiniRouter}
	martinInstance.Run()
}
