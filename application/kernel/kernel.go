package kernel

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/kernel/authority"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/application/kernel/router"
	"muidea.com/magicCenter/application/kernel/session"

	"log"

	"github.com/go-martini/martini"
)

// Kernel MagicCenter系统接口
type Kernel interface {
	// StartUp 启动系统
	StartUp() error
	// Run 运行系统
	Run()
	// ShutDown 关闭系统
	ShutDown()

	// ModuleHub 模块管理器
	ModuleHub() common.ModuleHub
	// Configuration 配置管理器
	Configuration() common.Configuration
}

type impl struct {
	loaderImpl          common.ModuleLoader
	configurationImpl   common.Configuration
	routerImpl          router.Router
	moduleHubImpl       common.ModuleHub
	authorityImpl       authority.Authority
	sessionRegistryImpl common.SessionRegistry
	instanceFrameImpl   *martini.Martini
}

// NewKernel 新建Kernel对象
func NewKernel(loader common.ModuleLoader, configuration common.Configuration) Kernel {
	i := &impl{
		loaderImpl:          loader,
		configurationImpl:   configuration,
		routerImpl:          router.CreateRouter(),
		moduleHubImpl:       modulehub.CreateModuleHub(),
		authorityImpl:       authority.CreateAuthority(),
		sessionRegistryImpl: session.CreateSessionRegistry(),
		instanceFrameImpl:   martini.New()}

	return i
}

func (i *impl) StartUp() error {
	i.loaderImpl.LoadAllModules(i.configurationImpl, i.sessionRegistryImpl, i.moduleHubImpl)

	allModules := i.moduleHubImpl.QueryAllModule()
	for _, m := range allModules {
		baseURL := m.URL()
		routes := m.Routes()
		if len(routes) > 0 {
			log.Printf("...............register %s's routes...............", m.Name())
		}
		for _, rt := range routes {
			i.routerImpl.AddRoute(baseURL, rt)
		}
	}

	i.moduleHubImpl.StartupAllModules()
	return nil
}

func (i *impl) Run() {
	martiniRouter := i.routerImpl.Router()

	i.instanceFrameImpl.Use(martini.Logger())
	i.instanceFrameImpl.Use(martini.Recovery())
	i.instanceFrameImpl.Use(authority.AuthorityHandler(i.authorityImpl))

	i.instanceFrameImpl.MapTo(martiniRouter, (*martini.Routes)(nil))
	i.instanceFrameImpl.Action(martiniRouter.Handle)

	instance := martini.ClassicMartini{}
	instance.Martini = i.instanceFrameImpl
	instance.Router = martiniRouter

	instance.Run()
}

func (i *impl) ShutDown() {
	/*
		退出时不需要做路由清理操作
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
	return i.sessionRegistryImpl.GetSession(w, r)
}
