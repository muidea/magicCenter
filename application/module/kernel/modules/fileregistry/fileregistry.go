package fileregistry

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/fileregistry/def"
	"muidea.com/magicCenter/application/module/kernel/modules/fileregistry/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/fileregistry/route"
)

type fileRegistry struct {
	routes              []common.Route
	fileRegistryHandler handler.FileRegistryHandler
}

// LoadModule 加载Static模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {
	instance := &fileRegistry{fileRegistryHandler: handler.CreateHandler()}

	instance.routes = route.AppendFileRegistryRoute(instance.routes)

	modHub.RegisterModule(instance)
}

// ID ID
func (instance *fileRegistry) ID() string {
	return def.ID
}

// Name 名称
func (instance *fileRegistry) Name() string {
	return def.Name
}

// Description 名称
func (instance *fileRegistry) Description() string {
	return def.Description
}

func (instance *fileRegistry) Group() string {
	return "resource"
}

func (instance *fileRegistry) Type() int {
	return common.INTERNAL
}

func (instance *fileRegistry) Status() int {
	return 0
}

func (instance *fileRegistry) EndPoint() interface{} {
	return nil
}

func (instance *fileRegistry) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route 路由信息
func (instance *fileRegistry) Routes() []common.Route {
	return instance.routes
}

// Startup 启动模块
func (instance *fileRegistry) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *fileRegistry) Cleanup() {

}
