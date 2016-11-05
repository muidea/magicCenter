package system

import (
	"magiccenter/common"
	"magiccenter/system/dbhelper"
	"magiccenter/system/modulehub"
	"magiccenter/system/router"
	"magiccenter/system/session"
	"martini"
	"net/http"
)

var routerImpl = router.CreateRouter()
var moduleHubImpl = modulehub.CreateModuleHub()
var instanceFrame *martini.Martini
var authImpl common.Authority
var configurationImpl common.Configuration

// GetRouter 获取系统的Router
func GetRouter() common.Router {
	return routerImpl
}

// GetModuleHub 获取系统的ModuleHub
func GetModuleHub() common.ModuleHub {
	return moduleHubImpl
}

// GetDBHelper 获取系统的数据库访问助手
func GetDBHelper() (common.DBHelper, error) {
	return dbhelper.NewHelper()
}

// GetSession 获取当前Session
func GetSession(w http.ResponseWriter, r *http.Request) common.Session {
	return session.GetSession(w, r)
}

// GetAuthority 获取当前Authority
func GetAuthority() common.Authority {
	return authImpl
}

// GetConfiguration 获取当前Configuration
func GetConfiguration() common.Configuration {
	return configurationImpl
}

// BindStatic 绑定静态资源路径
func BindStatic() {
	path := "template"
	path, _ = configurationImpl.GetOption(common.ResourcePath)
	instanceFrame.Use(martini.Static(path))
}

// BindAuthVerify 绑定权限校验器
func BindAuthVerify(auth common.Authority) {
	instanceFrame.Use(auth.Authority())
}

// Initialize 初始化
func Initialize(loader common.ModuleLoader, auth common.Authority, configuration common.Configuration) {
	instanceFrame = martini.New()
	authImpl = auth
	configurationImpl = configuration

	configurationImpl.LoadConfig()

	loader.LoadAllModules()

	allModules := moduleHubImpl.QueryAllModule()
	for _, m := range allModules {
		baseURL := m.URL()
		routes := m.Routes()
		for _, rt := range routes {
			routerImpl.AddRoute(baseURL, rt)
		}
	}

	BindStatic()

	BindAuthVerify(auth)

	moduleHubImpl.StartupAllModules()
}

// Uninitialize 反初始化
func Uninitialize() {

	allModules := moduleHubImpl.QueryAllModule()
	for _, m := range allModules {
		baseURL := m.URL()
		routes := m.Routes()
		for _, rt := range routes {
			routerImpl.RemoveRoute(baseURL, rt)
		}
	}

	moduleHubImpl.CleanupAllModules()
}

// Run 开始运行
func Run() {
	martiniRouter := routerImpl.Router()

	instanceFrame.Use(martini.Logger())
	instanceFrame.Use(martini.Recovery())
	instanceFrame.MapTo(martiniRouter, (*martini.Routes)(nil))
	instanceFrame.Action(martiniRouter.Handle)

	instance := martini.ClassicMartini{}
	instance.Martini = instanceFrame
	instance.Router = martiniRouter

	instance.Run()
}
