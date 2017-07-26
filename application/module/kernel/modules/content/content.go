package content

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/def"
	"muidea.com/magicCenter/application/module/kernel/modules/content/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/content/route"
)

type content struct {
	routes         []common.Route
	contentHandler common.ContentHandler
}

// LoadModule 加载模块
func LoadModule(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) {
	instance := &content{contentHandler: handler.CreateContentHandler()}

	instance.routes = route.AppendArticleRoute(instance.routes, instance.contentHandler, sessionRegistry)
	instance.routes = route.AppendCatalogRoute(instance.routes, instance.contentHandler, sessionRegistry)
	instance.routes = route.AppendLinkRoute(instance.routes, instance.contentHandler, sessionRegistry)
	instance.routes = route.AppendMediaRoute(instance.routes, instance.contentHandler, sessionRegistry)

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

func (instance *content) EntryPoint() interface{} {
	return instance.contentHandler
}

func (instance *content) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	groups = append(groups, model.AuthGroup{"PublicGroup", "允许查看公开权限的内容"})
	groups = append(groups, model.AuthGroup{"UserGroup", "允许查看用户权限范围内的内容"})
	groups = append(groups, model.AuthGroup{"AdminGroup", "允许管理用户权限范围内的内容"})

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
