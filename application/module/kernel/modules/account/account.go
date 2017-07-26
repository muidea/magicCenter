package account

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/account/def"
	"muidea.com/magicCenter/application/module/kernel/modules/account/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/account/route"
)

// LoadModule 加载模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {

	instance := &account{routes: make([]common.Route, 0), accountHandler: handler.CreateAccountHandler()}

	instance.routes = route.AppendUserRoute(instance.routes, instance.accountHandler)
	instance.routes = route.AppendGroupRoute(instance.routes, instance.accountHandler)

	modHub.RegisterModule(instance)
}

type account struct {
	routes         []common.Route
	accountHandler common.AccountHandler
}

func (instance *account) ID() string {
	return def.ID
}

func (instance *account) Name() string {
	return def.Name
}

func (instance *account) Description() string {
	return def.Description
}

func (instance *account) Group() string {
	return "kernel"
}

func (instance *account) Type() int {
	return common.KERNEL
}

func (instance *account) Status() int {
	return 0
}

func (instance *account) EntryPoint() interface{} {
	return instance.accountHandler
}

func (instance *account) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	groups = append(groups, model.AuthGroup{Name: "PublicGroup", Description: "允许查看公开权限的内容"})
	groups = append(groups, model.AuthGroup{Name: "UserGroup", Description: "允许查看用户权限范围内的内容"})
	groups = append(groups, model.AuthGroup{Name: "AdminGroup", Description: "允许管理用户权限范围内的内容"})

	return groups
}

// Route Account 路由信息
func (instance *account) Routes() []common.Route {
	return instance.routes
}

// Startup 启动Account模块
func (instance *account) Startup() bool {
	return true
}

// Cleanup 清除Account模块
func (instance *account) Cleanup() {

}
