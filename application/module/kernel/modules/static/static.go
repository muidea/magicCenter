package static

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/static/def"
	"muidea.com/magicCenter/application/module/kernel/modules/static/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/static/route"
)

type static struct {
	routes        []common.Route
	staticHandler common.StaticHandler
}

// LoadModule 加载Static模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {
	instance := &static{staticHandler: handler.CreateStaticHandler("./static/")}

	rt := route.CreateStaticViewRoute(instance.staticHandler)
	instance.routes = append(instance.routes, rt)

	rt = route.CreateStaticResRoute(instance.staticHandler)
	instance.routes = append(instance.routes, rt)

	modHub.RegisterModule(instance)
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
	return common.INTERNAL
}

func (instance *static) Status() int {
	return 0
}

func (instance *static) EndPoint() interface{} {
	return nil
}

func (instance *static) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
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
