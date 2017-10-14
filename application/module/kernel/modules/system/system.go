package system

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/system/def"
	"muidea.com/magicCenter/application/module/kernel/modules/system/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/system/route"
)

type system struct {
	routes        []common.Route
	systemHandler common.SystemHandler
}

// LoadModule 加载System模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &system{systemHandler: handler.CreateSystemHandler(configuration, sessionRegistry, moduleHub)}

	rt := route.CreateSystemRoute(instance.systemHandler)
	instance.routes = append(instance.routes, rt)

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
	return common.INTERNAL
}

func (instance *system) Status() int {
	return 0
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
