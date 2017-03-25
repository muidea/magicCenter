package authority

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/authority/handler"
	"muidea.com/magicCenter/application/module/kernel/authority/route"
)

// ID 模块ID
const ID = common.AuthorityModuleID

// Name 模块名称
const Name = "Magic Authority"

// Description 模块描述信息
const Description = "Magic 权限管理模块"

// URL 模块Url
const URL string = "/authority"

type authority struct {
	moduleHub        common.ModuleHub
	sessionRegistry  common.SessionRegistry
	routes           []common.Route
	authorityHandler common.AuthorityHandler
}

// LoadModule 加载模块
func LoadModule(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) {
	instance := &authority{
		moduleHub:        modHub,
		sessionRegistry:  sessionRegistry,
		authorityHandler: handler.CreateAuthorityHandler(modHub, sessionRegistry)}

	rt, _ := route.CreateAccountLoginRoute(instance.authorityHandler, sessionRegistry)
	instance.routes = append(instance.routes, rt)

	rt, _ = route.CreateAccountLogoutRoute(instance.authorityHandler, sessionRegistry)
	instance.routes = append(instance.routes, rt)

	modHub.RegisterModule(instance)
}

func (instance *authority) ID() string {
	return ID
}

func (instance *authority) Name() string {
	return Name
}

func (instance *authority) Description() string {
	return Description
}

func (instance *authority) Group() string {
	return "kernel"
}

func (instance *authority) Type() int {
	return common.KERNEL
}

func (instance *authority) URL() string {
	return URL
}

func (instance *authority) Status() int {
	return 0
}

func (instance *authority) EndPoint() interface{} {
	return instance.authorityHandler
}

func (instance *authority) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	groups = append(groups, model.CreateAuthGroup("PublicGroup", "允许查看公开权限的内容", ID))
	groups = append(groups, model.CreateAuthGroup("UserGroup", "允许查看用户权限范围内的内容", ID))
	groups = append(groups, model.CreateAuthGroup("AdminGroup", "允许管理用户权限范围内的内容", ID))

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
