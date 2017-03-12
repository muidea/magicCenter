package account

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/application/module/kernel/modules/account/handler"
)

// ID 模块ID
const ID = "b9e35167-b2a3-43ae-8c57-9b4379475e47"

// Name 模块名称
const Name = "Magic Account"

// Description 模块描述信息
const Description = "Magic 账号管理模块"

// URL 模块Url
const URL string = "/account"

// LoadModule 加载模块
func LoadModule(cfg configuration.Configuration, modHub modulehub.ModuleHub) {
	instance := &account{accountHandler: handler.CreateAccountHandler()}

	modHub.RegisterModule(instance)
}

type account struct {
	routes         []common.Route
	accountHandler common.AccountHandler
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

func (instance *account) Status() int {
	return 0
}

func (instance *account) EndPoint() interface{} {
	return instance.accountHandler
}

func (instance *account) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route Account 路由信息
func (instance *account) Routes() []common.Route {
	/*
		routes := []common.Route{
			// 用户账号信息管理视图
			common.NewRoute(common.GET, "manageUserView/", ui.ManageUserViewHandler),

			// 用户分组信息管理视图
			common.NewRoute(common.GET, "manageGroupView/", ui.ManageGroupViewHandler),
		}*/

	return instance.routes
}

// Startup 启动Account模块
func (instance *account) Startup() bool {
	return true
}

// Cleanup 清除Account模块
func (instance *account) Cleanup() {

}
