package content

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/def"
	"muidea.com/magicCenter/application/module/kernel/modules/content/handler"
)

type content struct {
	routes         []common.Route
	contentHandler common.ContentHandler
}

// LoadModule 加载模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {
	instance := &content{contentHandler: handler.CreateContentHandler()}

	modHub.RegisterModule(instance)
}

func (instance *content) ID() string {
	return def.ID
}

func (instance *content) Name() string {
	return def.Name
}

func (instance *content) Description() string {
	return def.Description
}

func (instance *content) Group() string {
	return "kernel"
}

func (instance *content) Type() int {
	return common.KERNEL
}

func (instance *content) Status() int {
	return 0
}

func (instance *content) EndPoint() interface{} {
	return instance.contentHandler
}

func (instance *content) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route 路由信息
func (instance *content) Routes() []common.Route {
	return instance.routes
}

// Startup 启动Content模块
func (instance *content) Startup() bool {
	return true
}

// Cleanup 清除Content模块
func (instance *content) Cleanup() {

}
