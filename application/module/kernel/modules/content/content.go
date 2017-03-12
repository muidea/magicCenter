package content

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/application/module/kernel/modules/content/handler"
)

// ID 模块ID
const ID = "3a7123ec-63f0-5e46-1234-e6ca1af6fe4e"

// Name 模块名称
const Name = "Magic Content"

// Description 模块描述信息
const Description = "Magic 内容管理模块"

// URL 模块Url
const URL string = "/content"

type content struct {
	routes         []common.Route
	contentHandler common.ContentHandler
}

// LoadModule 加载模块
func LoadModule(cfg configuration.Configuration, modHub modulehub.ModuleHub) {
	instance := &content{contentHandler: handler.CreateContentHandler()}

	modHub.RegisterModule(instance)
}

func (instance *content) ID() string {
	return ID
}

func (instance *content) Name() string {
	return Name
}

func (instance *content) Description() string {
	return Description
}

func (instance *content) Group() string {
	return "kernel"
}

func (instance *content) Type() int {
	return common.KERNEL
}

func (instance *content) URL() string {
	return URL
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
	/*
		routes := []common.Route{
			// Article管理视图
			router.NewRoute(common.GET, "manageArticleView/", ui.ManageArticleViewHandler, auth.AdminAuthVerify()),
			// Catalog管理视图
			router.NewRoute(common.GET, "manageCatalogView/", ui.ManageCatalogViewHandler, auth.AdminAuthVerify()),
			// Link管理视图
			router.NewRoute(common.GET, "manageLinkView/", ui.ManageLinkViewHandler, auth.AdminAuthVerify()),
			// Media管理视图
			router.NewRoute(common.GET, "manageMediaView/", ui.ManageMediaViewHandler, auth.AdminAuthVerify()),
		}
	*/
	return instance.routes
}

// Startup 启动Content模块
func (instance *content) Startup() bool {
	return true
}

// Cleanup 清除Content模块
func (instance *content) Cleanup() {

}
