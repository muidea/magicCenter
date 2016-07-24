package kernel

import (
	"log"
	"magiccenter/configuration"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/dashboard"
	"magiccenter/module"
	"magiccenter/modules/loader"
	"magiccenter/router"
	"martini"
)

var instanceFrame = martini.New()

// BindStatic 绑定静态资源路径
func BindStatic(path string) {
	instanceFrame.Use(martini.Static(path))
}

// BindAuthVerify 绑定权限校验器
func BindAuthVerify() {
	instanceFrame.Use(auth.Authority())
}

// Initialize 初始化Kernel
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

	BindAuthVerify()

	dashboard.RegisterRouter()

	loader.LoadAllModules()

	module.StartupAllModules()
}

// Uninitialize 反初始化Kernel，清除相关资源
func Uninitialize() {

	module.CleanupAllModules()
}

// Run 运行Kernel，进行路由分发
func Run() {
	martiniRouter := router.Router()

	instanceFrame.Use(martini.Logger())
	instanceFrame.Use(martini.Recovery())
	instanceFrame.MapTo(martiniRouter, (*martini.Routes)(nil))
	instanceFrame.Action(martiniRouter.Handle)

	instance := martini.ClassicMartini{}
	instance.Martini = instanceFrame
	instance.Router = martiniRouter

	instance.Run()
}
