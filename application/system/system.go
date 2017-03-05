package system

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/system/modulehub"
	"muidea.com/magicCenter/application/system/router"
	"muidea.com/magicCenter/application/system/session"

	"github.com/go-martini/martini"
)

type impl struct {
	loaderImpl        common.ModuleLoader
	configurationImpl common.Configuration
	routerImpl        common.Router
	moduleHubImpl     common.ModuleHub
	instanceFrame     *martini.Martini
}

// NewSystem 新建System对象
func NewSystem(loader common.ModuleLoader, configuration common.Configuration) common.System {
	i := &impl{
		loaderImpl:        loader,
		configurationImpl: configuration,
		routerImpl:        router.CreateRouter(),
		moduleHubImpl:     modulehub.CreateModuleHub(),
		instanceFrame:     martini.New()}

	return i
}

func (i *impl) StartUp() error {
	i.configurationImpl.LoadConfig()

	i.loaderImpl.LoadAllModules()

	allModules := i.moduleHubImpl.QueryAllModule()
	for _, m := range allModules {
		baseURL := m.URL()
		routes := m.Routes()
		for _, rt := range routes {
			i.routerImpl.AddRoute(baseURL, rt)
		}
	}

	i.moduleHubImpl.StartupAllModules()
	return nil
}

func (i *impl) Run() {
	martiniRouter := i.routerImpl.Router()

	i.instanceFrame.Use(martini.Logger())
	i.instanceFrame.Use(martini.Recovery())
	i.instanceFrame.MapTo(martiniRouter, (*martini.Routes)(nil))
	i.instanceFrame.Action(martiniRouter.Handle)

	instance := martini.ClassicMartini{}
	instance.Martini = i.instanceFrame
	instance.Router = martiniRouter

	instance.Run()
}

func (i *impl) ShutDown() error {
	/* 退出时不需要做路由清理操作
	allModules := i.moduleHubImpl.QueryAllModule()
	for _, m := range allModules {
		baseURL := m.URL()
		routes := m.Routes()
		for _, rt := range routes {
			routerImpl.RemoveRoute(baseURL, rt)
		}
	}
	*/

	i.moduleHubImpl.CleanupAllModules()
	return nil
}

// GetRouter 获取系统的Router
func (i *impl) Router() common.Router {
	return i.routerImpl
}

// GetModuleHub 获取系统的ModuleHub
func (i *impl) ModuleHub() common.ModuleHub {
	return i.moduleHubImpl
}

// GetConfiguration 获取当前Configuration
func (i *impl) Configuration() common.Configuration {
	return i.configurationImpl
}

// GetSession 获取当前Session
func (i *impl) Session(w http.ResponseWriter, r *http.Request) common.Session {
	return session.GetSession(w, r)
}
