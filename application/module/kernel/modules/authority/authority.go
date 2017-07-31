package authority

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/def"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/route"
)

type authority struct {
	moduleHub        common.ModuleHub
	sessionRegistry  common.SessionRegistry
	routes           []common.Route
	authorityHandler common.AuthorityHandler
}

// LoadModule 加载模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &authority{
		moduleHub:        moduleHub,
		sessionRegistry:  sessionRegistry,
		authorityHandler: handler.CreateAuthorityHandler(moduleHub, sessionRegistry)}

	instance.routes = route.AppendACLRoute(instance.routes, instance.authorityHandler, sessionRegistry)
	instance.routes = route.AppendAuthGropRoute(instance.routes, instance.authorityHandler, sessionRegistry)

	moduleHub.RegisterModule(instance)
}

func (instance *authority) ID() string {
	return def.ID
}

func (instance *authority) Name() string {
	return def.Name
}

func (instance *authority) Description() string {
	return def.Description
}

func (instance *authority) Group() string {
	return "kernel"
}

func (instance *authority) Type() int {
	return common.KERNEL
}

func (instance *authority) Status() int {
	return 0
}

func (instance *authority) EntryPoint() interface{} {
	return instance.authorityHandler
}

func (instance *authority) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	groups = append(groups, model.AuthGroup{Name: "PublicGroup", Description: "允许查看公开权限的内容"})
	groups = append(groups, model.AuthGroup{Name: "UserGroup", Description: "允许查看用户权限范围内的内容"})
	groups = append(groups, model.AuthGroup{Name: "AdminGroup", Description: "允许管理用户权限范围内的内容"})

	return groups
}

// Route 路由信息
func (instance *authority) Routes() []common.Route {
	return instance.routes
}

// Startup 启动模块
func (instance *authority) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *authority) Cleanup() {

}
