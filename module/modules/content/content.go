package content

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/content/def"
	"muidea.com/magicCenter/module/modules/content/handler"
	"muidea.com/magicCenter/module/modules/content/route"
	common_const "muidea.com/magicCommon/common"
)

type content struct {
	routes         []common.Route
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

// LoadModule 加载模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	accountModule, _ := moduleHub.FindModule(common.AccountModuleID)
	accountHandler := accountModule.EntryPoint().(common.AccountHandler)

	fileRegistryModule, _ := moduleHub.FindModule(common.FileRegistryModuleID)
	fileRegistryHandler := fileRegistryModule.EntryPoint().(common.FileRegistryHandler)

	instance := &content{contentHandler: handler.CreateContentHandler(), accountHandler: accountHandler}

	instance.routes = route.AppendSummaryRoute(instance.routes, instance.contentHandler, instance.accountHandler, sessionRegistry)
	instance.routes = route.AppendArticleRoute(instance.routes, instance.contentHandler, instance.accountHandler, sessionRegistry)
	instance.routes = route.AppendCatalogRoute(instance.routes, instance.contentHandler, instance.accountHandler, sessionRegistry)
	instance.routes = route.AppendLinkRoute(instance.routes, instance.contentHandler, instance.accountHandler, sessionRegistry)
	instance.routes = route.AppendMediaRoute(instance.routes, instance.contentHandler, instance.accountHandler, fileRegistryHandler, sessionRegistry)

	moduleHub.RegisterModule(instance)
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
	return common_const.KERNEL
}

func (instance *content) Status() int {
	return common_const.ACTIVE
}

func (instance *content) EntryPoint() interface{} {
	return instance.contentHandler
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
