package kernel

import (
	"log"
	"martini"
	"magiccenter/module"
	"magiccenter/router"
	"magiccenter/configuration"
	"magiccenter/kernel/admin"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/account"
	"magiccenter/kernel/system"
	"magiccenter/kernel/content"
	"magiccenter/modules/loader"	
)

var instanceFrame = martini.New()

func BindStatic(path string) {
	instanceFrame.Use(martini.Static(path))
}

func Initialize() {
	log.Println("initialize kernel...")
	
	configuration.LoadConfig()
	
	staticPath, found := configuration.GetOption(configuration.STATIC_PATH)
	if found {
		BindStatic(staticPath)
	}
	
	resourceFilePath, found := configuration.GetOption(configuration.RESOURCE_PATH)
	if found {
		BindStatic(resourceFilePath)
	}
	
	admin.RegisterRouter()
	auth.RegisterRouter()
	account.RegisterRouter()
	system.RegisterRouter()
	content.RegisterRouter()
	
	loader.LoadAllModules()
	
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
