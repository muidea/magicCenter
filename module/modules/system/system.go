package system

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/system/def"
	"muidea.com/magicCenter/module/modules/system/handler"
	"muidea.com/magicCenter/module/modules/system/route"
	common_const "muidea.com/magicCommon/common"
)

type system struct {
	routes        []common.Route
	systemHandler common.SystemHandler
}

// LoadModule 加载System模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &system{systemHandler: handler.CreateSystemHandler(configuration, sessionRegistry, moduleHub)}

	instance.routes = route.AppendSystemRoute(instance.routes, instance.systemHandler)

	moduleHub.RegisterModule(instance)
}

// ID System ID
func (instance *system) ID() string {
	return def.ID
}

// Name System 名称
func (instance *system) Name() string {
	return def.Name
}

// Description System名称
func (instance *system) Description() string {
	return def.Description
}

func (instance *system) Group() string {
	return "resource"
}

func (instance *system) Type() int {
	return common_const.INTERNAL
}

func (instance *system) Status() int {
	return common_const.ACTIVE
}

func (instance *system) EntryPoint() interface{} {
	return instance.systemHandler
}

// Route System 路由信息
func (instance *system) Routes() []common.Route {
	return instance.routes
}

// Startup 启动System模块
func (instance *system) Startup() bool {
	return true
}

// Cleanup 清除System模块
func (instance *system) Cleanup() {

}
