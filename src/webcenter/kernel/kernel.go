package kernel

import (
	"log"
	"martini"
	"webcenter/module"
)

var instanceRouter = NewRouter()
var instanceFrame = martini.New()

func RegisterGetHandler(pattern string, h interface{}) {
	instanceRouter.AddGetRoute(pattern, h)	
}

func UnRegisterGetHandler(pattern string, h interface{}) {
}

func RegisterPostHandler(pattern string, h interface{}) {
	instanceRouter.AddPostRoute(pattern, h)
}

func UnRegisterPostHandler(pattern string, h interface{}) {
}

func BindStatic(path string) {
	instanceFrame.Use(martini.Static(path))
}

func Initialize() {
	log.Println("initialize kernel...")
	
	BindStatic(staticPath)
	BindStatic(resourceFilePath)	
	module.StartupAllModules(instanceRouter)
}


func Uninitialize() {
	module.CleanupAllModules()
}

func Run() {
	martiniRouter := instanceRouter.Router()
	
	instanceFrame.Use(martini.Logger())
	instanceFrame.Use(martini.Recovery())
	instanceFrame.MapTo(martiniRouter, (*martini.Routes)(nil))
	instanceFrame.Action(martiniRouter.Handle)

	martinInstance := &martini.ClassicMartini{instanceFrame, martiniRouter}
	martinInstance.Run()
}