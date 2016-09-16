package dashboard

import (
	"magiccenter/common"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/modules/dashboard/ui"
	"magiccenter/module"
	"magiccenter/router"
)

// ID 模块ID
const ID = "f67123ea-6fe0-5e46-1234-e6ca1af6fe4e"

// Name 模块名称
const Name = "Magic Dashboard"

// Description 模块描述信息
const Description = "Magic Dashboard模块"

// URL 模块Url
const URL string = "admin"

type dashboard struct {
}

var instance *dashboard

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &dashboard{}
	}

	module.RegisterModule(instance)
}

func (instance *dashboard) ID() string {
	return ID
}

func (instance *dashboard) Name() string {
	return Name
}

func (instance *dashboard) Description() string {
	return Description
}

func (instance *dashboard) Group() string {
	return "admin dashboard"
}

func (instance *dashboard) Type() int {
	return common.KERNEL
}

func (instance *dashboard) URL() string {
	return URL
}

func (instance *dashboard) EndPoint() common.EndPoint {
	return nil
}

// Route 路由信息
func (instance *dashboard) Routes() []common.Route {
	routes := []common.Route{
		// 管理视图
		router.NewRoute(common.GET, "/", ui.AdminViewHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.GET, "login/", ui.LoginViewHandler, nil),
		router.NewRoute(common.POST, "verify/", ui.VerifyAuthActionHandler, nil),
		router.NewRoute(common.GET, "logout/", ui.LogoutActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *dashboard) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *dashboard) Cleanup() {

}

// Invoke 执行外部命令
func (instance *dashboard) Invoke(param interface{}, result interface{}) bool {
	return false
}