package endpoint

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/endpoint/def"
	"muidea.com/magicCenter/module/modules/endpoint/handler"
	"muidea.com/magicCenter/module/modules/endpoint/route"
	common_const "muidea.com/magicCommon/common"
)

type endpoint struct {
	moduleHub       common.ModuleHub
	sessionRegistry common.SessionRegistry
	routes          []common.Route
	endpointHandler common.EndpointHandler
	accountHandler  common.AccountHandler
}

// LoadModule 加载模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	accountModule, _ := moduleHub.FindModule(common.AccountModuleID)
	accountHandler := accountModule.EntryPoint().(common.AccountHandler)

	instance := &endpoint{
		moduleHub:       moduleHub,
		sessionRegistry: sessionRegistry,
		accountHandler:  accountHandler,
		endpointHandler: handler.CreateEndpointHandler()}

	instance.routes = route.AppendEndpointRoute(instance.routes, instance.endpointHandler, instance.accountHandler, moduleHub, sessionRegistry)

	moduleHub.RegisterModule(instance)
}

func (instance *endpoint) ID() string {
	return def.ID
}

func (instance *endpoint) Name() string {
	return def.Name
}

func (instance *endpoint) Description() string {
	return def.Description
}

func (instance *endpoint) Group() string {
	return "kernel"
}

func (instance *endpoint) Type() int {
	return common_const.KERNEL
}

func (instance *endpoint) Status() int {
	return common_const.ACTIVE
}

func (instance *endpoint) EntryPoint() interface{} {
	return instance.endpointHandler
}

// Route 路由信息
func (instance *endpoint) Routes() []common.Route {
	return instance.routes
}

// Startup 启动模块
func (instance *endpoint) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *endpoint) Cleanup() {

}
