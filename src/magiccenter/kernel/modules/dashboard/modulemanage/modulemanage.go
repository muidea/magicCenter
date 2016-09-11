package modulemanage

import (
	"magiccenter/common"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/modules/dashboard/modulemanage/ui"
	"magiccenter/module"
	"magiccenter/router"
)

// ID 模块ID
const ID = "6a384b24-8fb1-4e28-885c-c3aa9480ae0c"

// Name 模块名称
const Name = "Magic ModuleManage"

// Description 模块描述信息
const Description = "Magic 模块管理模块"

// URL 模块Url
const URL string = "admin"

type modulemanage struct {
}

var instance *modulemanage

// LoadModule 加载ModuleManage模块
func LoadModule() {
	if instance == nil {
		instance = &modulemanage{}
	}

	module.RegisterModule(instance)
}

func (instance *modulemanage) ID() string {
	return ID
}

func (instance *modulemanage) Name() string {
	return Name
}

func (instance *modulemanage) Description() string {
	return Description
}

func (instance *modulemanage) Group() string {
	return "admin modulemanage"
}

func (instance *modulemanage) Type() int {
	return common.KERNEL
}

func (instance *modulemanage) URL() string {
	return URL
}

func (instance *modulemanage) EndPoint() common.EndPoint {
	return nil
}

// Route 路由信息
func (instance *modulemanage) Routes() []common.Route {
	routes := []common.Route{
		// 用户账号信息管理视图
		router.NewRoute(common.GET, "modulemanage/", ui.ModuleManageViewHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *modulemanage) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *modulemanage) Cleanup() {

}

// Invoke 执行外部命令
func (instance *modulemanage) Invoke(param interface{}) bool {
	return false
}
