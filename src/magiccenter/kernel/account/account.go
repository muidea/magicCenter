package account

import (
	"magiccenter/kernel/account/ui"
	"magiccenter/kernel/auth"
	"magiccenter/module"
	"magiccenter/router"
)

// ID Account模块ID
const ID = "f67123ec-63f0-5e46-0000-e6ca1af6fe4e"

// Name Account模块名称
const Name = "Magic Account"

// Description Account模块描述信息
const Description = "Magic 账号管理模块"

// URL Account模块Url
const URL string = "account"

type account struct {
}

var instance *account

// LoadModule 加载Account模块
func LoadModule() {
	if instance == nil {
		instance = &account{}
	}

	module.RegisterModule(instance)
}

// ID Account ID
func (instance *account) ID() string {
	return ID
}

// Name Account 名称
func (instance *account) Name() string {
	return Name
}

// Description Account名称
func (instance *account) Description() string {
	return Description
}

func (instance *account) Group() string {
	return "kernel"
}

func (instance *account) Type() int {
	return module.KERNEL
}

// URL Account url
func (instance *account) URL() string {
	return URL
}

func (instance *account) Resource() module.Resource {
	return nil
}

// Route Account 路由信息
func (instance *account) Routes() []router.Route {
	routes := []router.Route{
		// 用户账号信息管理
		router.NewRoute(router.GET, "/admin/account/manageUser/", ui.ManageUserHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/account/queryAllUser/", ui.QueryAllUserHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/account/queryUser/", ui.QueryUserHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/account/deleteUser/", ui.DeleteUserHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/account/checkAccount/", ui.CheckAccountHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/account/ajaxUser/", ui.AjaxUserHandler, auth.AdminAuthVerify()),

		// 用户分组信息管理
		router.NewRoute(router.GET, "/admin/account/manageGroup/", ui.ManageGroupHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/account/queryAllGroup/", ui.QueryAllGroupHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/account/queryGroup/", ui.QueryGroupHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/account/deleteGroup/", ui.DeleteGroupHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/account/ajaxGroup/", ui.AjaxGroupHandler, auth.AdminAuthVerify()),

		// 用户空间管理
		router.NewRoute(router.GET, "/user/profile/", ui.UserProfileViewHandler, nil),
		router.NewRoute(router.GET, "/user/verify/", ui.UserVerifyViewHandler, nil),
		router.NewRoute(router.POST, "/user/ajaxVerify/", ui.AjaxVerifyHandler, nil),
	}

	return routes
}

// Startup 启动Account模块
func (instance *account) Startup() bool {
	return true
}

// Cleanup 清除Account模块
func (instance *account) Cleanup() {

}

// Invoke 执行外部命令
func (instance *account) Invoke(param interface{}) bool {
	return false
}
