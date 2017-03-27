package dashboard

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/dashboard/def"
	"muidea.com/magicCenter/application/module/kernel/modules/dashboard/route"
)

// 授权分组属性Key，用于读取和存储授权分组信息
const authGroupKey = "f67123ea-6fe0-5e46-1234-e6ca1af6fe4e_authGroupKey"

type dashboard struct {
	routes []common.Route
}

// LoadModule 加载模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {
	instance := &dashboard{}

	instance.routes = route.AppendModuleRoute(instance.routes, modHub)

	modHub.RegisterModule(instance)
}

func (instance *dashboard) ID() string {
	return def.ID
}

func (instance *dashboard) Name() string {
	return def.Name
}

func (instance *dashboard) Description() string {
	return def.Description
}

func (instance *dashboard) Group() string {
	return "admin dashboard"
}

func (instance *dashboard) Type() int {
	return common.KERNEL
}

func (instance *dashboard) Status() int {
	return 0
}

func (instance *dashboard) EndPoint() interface{} {
	return nil
}

func (instance *dashboard) AuthGroups() []model.AuthGroup {
	authGroup := []model.AuthGroup{}
	return authGroup
}

// Route 路由信息
func (instance *dashboard) Routes() []common.Route {
	return instance.routes
}

// Startup 启动模块
func (instance *dashboard) Startup() bool {

	return true
}

// Cleanup 清除模块
func (instance *dashboard) Cleanup() {

}
