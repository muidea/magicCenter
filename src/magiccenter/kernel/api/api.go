package api

import (
	"magiccenter/common"
	"magiccenter/kernel/api/ui"
	"magiccenter/system"

	"muidea.com/util"
)

// ID 模块ID
const ID = "5fa671dc-ccb5-4005-8500-f0e45b13705b"

// Name 模块名称
const Name = "Magic Dashboard API"

// Description 模块描述信息
const Description = "Magic Dashboard API模块"

// URL 模块Url
const URL string = "/api"

type api struct {
}

var instance *api

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &api{}
	}

	modulehub := system.GetModuleHub()
	modulehub.RegisterModule(instance)
}

func (instance *api) ID() string {
	return ID
}

func (instance *api) Name() string {
	return Name
}

func (instance *api) Description() string {
	return Description
}

func (instance *api) Group() string {
	return "admin api"
}

func (instance *api) Type() int {
	return common.KERNEL
}

func (instance *api) URL() string {
	return URL
}

func (instance *api) Status() int {
	return 0
}

func (instance *api) EndPoint() common.EndPoint {
	return nil
}

// Route 路由信息
func (instance *api) Routes() []common.Route {
	router := system.GetRouter()
	auth := system.GetAuthority()

	routes := []common.Route{
		// 获取Module列表
		router.NewRoute(common.GET, "module/", ui.GetModuleListActionHandler, auth.AdminAuthVerify()),
		// 获取ModuleBlock
		router.NewRoute(common.GET, "module/block/", ui.GetModuleBlockActionHandler, auth.AdminAuthVerify()),
		// 获取ModuleContent
		router.NewRoute(common.GET, "module/content/", ui.GetModuleContentActionHandler, auth.AdminAuthVerify()),
		// 获取ModuleContent
		router.NewRoute(common.GET, "module/authority/", ui.GetModuleAuthorityGroupActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *api) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *api) Cleanup() {

}

// Invoke 执行外部命令
func (instance *api) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}
