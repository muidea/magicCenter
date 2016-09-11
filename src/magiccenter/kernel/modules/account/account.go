package account

import (
	"magiccenter/common"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/modules/account/ui"
	"magiccenter/module"
	"magiccenter/router"
)

// ID 模块ID
const ID = "f67123ec-63f0-5e46-0000-e6ca1af6fe4e"

// Name 模块名称
const Name = "Magic Account"

// Description 模块描述信息
const Description = "Magic 账号管理模块"

// URL 模块Url
const URL string = "account"

type account struct {
}

var instance *account

// LoadModule 加载模块
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
	return common.KERNEL
}

func (instance *account) URL() string {
	return URL
}

func (instance *account) EndPoint() common.EndPoint {
	return nil
}

// Route Account 路由信息
func (instance *account) Routes() []common.Route {
	routes := []common.Route{
		// 用户账号信息管理视图
		router.NewRoute(common.GET, "manageUserView/", ui.ManageUserViewHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.GET, "queryAllUser/", ui.QueryAllUserActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.GET, "queryUser/", ui.QueryUserActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.GET, "deleteUser/", ui.DeleteUserActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.POST, "checkAccount/", ui.CheckAccountActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.POST, "ajaxUser/", ui.SaveUserActionHandler, auth.AdminAuthVerify()),

		// 用户分组信息管理视图
		router.NewRoute(common.GET, "manageGroupView/", ui.ManageGroupViewHandler, auth.AdminAuthVerify()),

		router.NewRoute(common.GET, "queryAllGroup/", ui.QueryAllGroupActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.GET, "queryGroup/", ui.QueryGroupActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.GET, "deleteGroup/", ui.DeleteGroupActionHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.POST, "ajaxGroup/", ui.SaveGroupActionHandler, auth.AdminAuthVerify()),
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
