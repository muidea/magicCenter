package kernel

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/kernel/authority"
	"muidea.com/magicCenter/kernel/modulehub"
	"muidea.com/magicCenter/kernel/router"
	"muidea.com/magicCenter/kernel/session"
	"muidea.com/magicCommon/model"
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

func (i *impl) isStartup() bool {
	_, ok := i.configurationImpl.GetOption(model.AppStartupData)
	return ok
}

func (i *impl) startuped() {
	startupFlag := fmt.Sprintf("startup_TimeStamp:%s", time.Now().Format("2006-01-02 15:04:05"))
	i.configurationImpl.SetOption(model.AppStartupData, startupFlag)
}

func (i *impl) StartUp() error {
	i.loaderImpl.LoadAllModules(i.configurationImpl, i.sessionRegistryImpl, i.moduleHubImpl)

	isStartup := i.isStartup()
	var authorityHandler common.AuthorityHandler
	if !isStartup {
		mod, ok := i.moduleHubImpl.FindModule(common.AuthorityModuleID)
		if ok {
			authorityHandler = mod.EntryPoint().(common.AuthorityHandler)
		}
	}

	allModules := i.moduleHubImpl.GetAllModule()
	for _, m := range allModules {
		routes := m.Routes()
		// if define routes trace something...
		if len(routes) > 0 {
			log.Printf("...............register %s's routes...............", m.Name())
		}
		for _, rt := range routes {
			i.routerImpl.AddRoute(rt)

			if !isStartup && authorityHandler != nil {
				authorityHandler.InsertACL(rt.Pattern(), rt.Method(), m.ID(), 0, rt.AuthGroup())
			}
		}
	}

	if !isStartup {
		i.startuped()
	}

	i.moduleHubImpl.StartupAllModules()
	return nil
}

func (i *impl) Run() {
	martiniRouter := i.routerImpl.Router()

	i.instanceFrameImpl.Use(martini.Logger())
	i.instanceFrameImpl.Use(martini.Recovery())
	i.instanceFrameImpl.Use(authority.VerifyHandler(i.moduleHubImpl, i.authorityImpl))

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
		allModules := i.moduleHubImpl.GetAllModule()
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
