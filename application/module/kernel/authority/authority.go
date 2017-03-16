package authority

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/authority/handler"
	"muidea.com/magicCenter/application/module/kernel/authority/route"
	"muidea.com/magicCenter/foundation/cache"
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
	cache := cache.NewCache()
	instance := &authority{
		moduleHub:        modHub,
		sessionRegistry:  sessionRegistry,
		authorityHandler: handler.CreateAuthorityHandler(modHub, cache)}

	rt, _ := route.CreateAccountLoginRoute(modHub, sessionRegistry)
	instance.routes = append(instance.routes, rt)

	rt, _ = route.CreateAccountLogoutRoute(modHub, sessionRegistry)
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

	return groups
}

// Route 路由信息
func (instance *authority) Routes() []common.Route {
	routes := []common.Route{}

	return routes
}

// Startup 启动模块
func (instance *authority) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *authority) Cleanup() {

}
