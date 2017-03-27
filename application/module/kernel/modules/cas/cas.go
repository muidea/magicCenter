package cas

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/cas/def"
	"muidea.com/magicCenter/application/module/kernel/modules/cas/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/cas/route"
)

type authority struct {
	moduleHub       common.ModuleHub
	sessionRegistry common.SessionRegistry
	routes          []common.Route
	casHandler      common.CASHandler
}

// LoadModule 加载模块
func LoadModule(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) {
	instance := &authority{
		moduleHub:       modHub,
		sessionRegistry: sessionRegistry,
		casHandler:      handler.CreateCASHandler(modHub, sessionRegistry)}

	rt, _ := route.CreateAccountLoginRoute(instance.casHandler, sessionRegistry)
	instance.routes = append(instance.routes, rt)

	rt, _ = route.CreateAccountLogoutRoute(instance.casHandler, sessionRegistry)
	instance.routes = append(instance.routes, rt)

	modHub.RegisterModule(instance)
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

func (instance *authority) EndPoint() interface{} {
	return instance.casHandler
}

func (instance *authority) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	groups = append(groups, model.CreateAuthGroup("PublicGroup", "允许查看公开权限的内容", def.ID))
	groups = append(groups, model.CreateAuthGroup("UserGroup", "允许查看用户权限范围内的内容", def.ID))
	groups = append(groups, model.CreateAuthGroup("AdminGroup", "允许管理用户权限范围内的内容", def.ID))

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
