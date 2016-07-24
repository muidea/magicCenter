package account

import (
	"magiccenter/kernel/auth"
	"magiccenter/kernel/modules/account/ui"
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

func (instance *account) ID() string {
	return ID
}

func (instance *account) Name() string {
	return Name
}

func (instance *account) Description() string {
	return Description
}

func (instance *account) Group() string {
	return "kernel"
}

func (instance *account) Type() int {
	return module.KERNEL
}

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
		router.NewRoute(router.GET, "manageUserView/", ui.ManageUserViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "queryAllUser/", ui.QueryAllUserActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "queryUser/", ui.QueryUserActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "deleteUser/", ui.DeleteUserActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "checkAccount/", ui.CheckAccountActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "ajaxUser/", ui.SaveUserActionHandler, auth.AdminAuthVerify()),

		// 用户分组信息管理
		router.NewRoute(router.GET, "manageGroup/", ui.ManageGroupViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "queryAllGroup/", ui.QueryAllGroupActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "queryGroup/", ui.QueryGroupActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "deleteGroup/", ui.DeleteGroupActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "ajaxGroup/", ui.SaveGroupActionHandler, auth.AdminAuthVerify()),
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
