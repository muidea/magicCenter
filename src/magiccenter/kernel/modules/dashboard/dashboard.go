package dashboard

import (
	"magiccenter/common"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/modules/dashboard/ui"
	"magiccenter/module"
	"magiccenter/router"
)

// ID Dashboard模块ID
const ID = "f67123ea-6fe0-5e46-1234-e6ca1af6fe4e"

// Name Account模块名称
const Name = "Magic Dashboard"

// Description Dashboard模块描述信息
const Description = "Magic Dashboard模块"

// URL Account模块Url
const URL string = "admin"

type dashboard struct {
}

var instance *dashboard

// LoadModule 加载Dashboard模块
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
	return "kernel"
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

// Route Account 路由信息
func (instance *dashboard) Routes() []common.Route {
	routes := []common.Route{
		// 用户账号信息管理视图
		router.NewRoute(common.GET, "/", ui.AdminViewHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.GET, "login/", ui.LoginViewHandler, nil),
		router.NewRoute(common.POST, "verify/", ui.VerifyAuthActionHandler, nil),
		router.NewRoute(common.GET, "logout/", ui.LogoutActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动Account模块
func (instance *dashboard) Startup() bool {
	return true
}

// Cleanup 清除Account模块
func (instance *dashboard) Cleanup() {

}

// Invoke 执行外部命令
func (instance *dashboard) Invoke(param interface{}) bool {
	return false
}
