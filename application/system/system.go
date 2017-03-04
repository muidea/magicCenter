package system

import (
	"net/http"
	"path"

	"muidea.com/magiccenter/application/system/dbhelper"
	"muidea.com/magiccenter/application/system/modulehub"
	"muidea.com/magiccenter/application/system/router"
	"muidea.com/magiccenter/application/system/session"

	"muidea.com/magiccenter/application/common"

	"github.com/go-martini/martini"
)

// System MagicCenter系统
type System interface {
	// Router 路由器
	Router() common.Router, error
	// ModuleHub 模块管理器
	ModuleHub() commmon.ModuleHub, error
	// Configuration 配置管理器
	Configuration() common.Configuration, error
	// DBHelper 数据库管理器
	DBHelper() common.DBHelper, error
	// Session 当前Session
	Session(w http.ResponseWriter, r *http.Request) common.Session
	// Authority 权限校验器
	Authority(w http.ResponseWriter, r *http.Request) common.Authority
}

// NewSystem 新建System对象
func NewSystem() System {
	i := &impl{}

	return i
}

type impl struct {
	routerImpl common.Router
	moduleHubImpl common.ModuleHub
	configurationImpl common.Configuration
	authImpl comon.Authority
}

var routerImpl = router.CreateRouter()
var moduleHubImpl = modulehub.CreateModuleHub()
var instanceFrame *martini.Martini
var authImpl common.Authority
var configurationImpl common.Configuration

// GetRouter 获取系统的Router
func (i *impl)Router() common.Router {
	return routerImpl
}

// GetModuleHub 获取系统的ModuleHub
func (i *impl)ModuleHub() common.ModuleHub {
	return moduleHubImpl
}

// GetDBHelper 获取系统的数据库访问助手
func (i *impl)DBHelper() (common.DBHelper, error) {
	return dbhelper.NewHelper()
}

// GetSession 获取当前Session
func (i *impl)Session(w http.ResponseWriter, r *http.Request) common.Session {
	return session.GetSession(w, r)
}

// GetAuthority 获取当前Authority
func (i *impl)Authority() common.Authority {
	return authImpl
}

// GetConfiguration 获取当前Configuration
func (i *impl)Configuration() common.Configuration {
	return configurationImpl
}

// bindResourcePath 绑定资源路径
func bindResourcePath() {
	path := "template/resources"
	instanceFrame.Use(martini.Static(path))
}

// bindAuthVerify 绑定权限校验器
func bindAuthVerify(auth common.Authority) {
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

	bindResourcePath()

	bindAuthVerify(auth)

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
