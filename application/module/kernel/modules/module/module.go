package module

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/module/def"
	"muidea.com/magicCenter/application/module/kernel/modules/module/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/module/route"
	common_const "muidea.com/magicCommon/common"
)

type module struct {
	routes        []common.Route
	moduleHandler common.ModuleRegistryHandler
}

// LoadModule 加载System模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &module{moduleHandler: handler.CreateModuleRegistryHandler(configuration, sessionRegistry, moduleHub)}

	instance.routes = route.AppendModuleRegistryRoute(instance.routes, instance.moduleHandler)

	moduleHub.RegisterModule(instance)
}

// ID System ID
func (instance *module) ID() string {
	return def.ID
}

// Name System 名称
func (instance *module) Name() string {
	return def.Name
}

// Description System名称
func (instance *module) Description() string {
	return def.Description
}

func (instance *module) Group() string {
	return "resource"
}

func (instance *module) Type() int {
	return common_const.INTERNAL
}

func (instance *module) Status() int {
	return common_const.ACTIVE
}

func (instance *module) EntryPoint() interface{} {
	return instance.moduleHandler
}

// Route System 路由信息
func (instance *module) Routes() []common.Route {
	return instance.routes
}

// Startup 启动System模块
func (instance *module) Startup() bool {
	return true
}

// Cleanup 清除System模块
func (instance *module) Cleanup() {

}
