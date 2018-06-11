package static

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/static/def"
	"muidea.com/magicCenter/module/modules/static/handler"
	"muidea.com/magicCenter/module/modules/static/route"
	common_const "muidea.com/magicCommon/common"
)

type static struct {
	routes        []common.Route
	staticHandler common.StaticHandler
}

// LoadModule 加载Static模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &static{staticHandler: handler.CreateStaticHandler(configuration, sessionRegistry, moduleHub)}

	rt := route.CreateStaticResRoute(instance.staticHandler)
	instance.routes = append(instance.routes, rt)

	moduleHub.RegisterModule(instance)
}

// ID Static ID
func (instance *static) ID() string {
	return def.ID
}

// Name Static 名称
func (instance *static) Name() string {
	return def.Name
}

// Description Static名称
func (instance *static) Description() string {
	return def.Description
}

func (instance *static) Group() string {
	return "resource"
}

func (instance *static) Type() int {
	return common_const.INTERNAL
}

func (instance *static) Status() int {
	return common_const.ACTIVE
}

func (instance *static) EntryPoint() interface{} {
	return nil
}

// Route Static 路由信息
func (instance *static) Routes() []common.Route {
	return instance.routes
}

// Startup 启动Static模块
func (instance *static) Startup() bool {
	return true
}

// Cleanup 清除Static模块
func (instance *static) Cleanup() {

}
