package content

import (
	"magiccenter/kernel/auth"
	"magiccenter/kernel/modules/content/ui"
	"magiccenter/module"
	"magiccenter/router"
)

// ID Content模块ID
const ID = "3a7123ec-63f0-5e46-1234-e6ca1af6fe4e"

// Name Content模块名称
const Name = "Magic Content"

// Description Content模块描述信息
const Description = "Magic 内容管理模块"

// URL Content模块Url
const URL string = "content"

type content struct {
}

var instance *content

// LoadModule 加载Content模块
func LoadModule() {
	if instance == nil {
		instance = &content{}
	}

	module.RegisterModule(instance)
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
	return module.KERNEL
}

func (instance *content) URL() string {
	return URL
}

func (instance *content) Resource() module.Resource {
	return nil
}

// Route Content 路由信息
func (instance *content) Routes() []router.Route {
	routes := []router.Route{
		// Article管理
		router.NewRoute(router.GET, "/admin/content/manageArticle/", ui.ManageArticleHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryAllArticle/", ui.QueryAllArticleHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryArticle/", ui.QueryArticleHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/deleteArticle/", ui.DeleteArticleHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/ajaxArticle/", ui.AjaxArticleHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/editArticle/", ui.EditArticleHandler, auth.AdminAuthVerify()),

		// Catalog管理
		router.NewRoute(router.GET, "/admin/content/manageCatalog/", ui.ManageCatalogHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryAllCatalog/", ui.QueryAllCatalogHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryCatalog/", ui.QueryCatalogHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/deleteCatalog/", ui.DeleteCatalogHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/ajaxCatalog/", ui.AjaxCatalogHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/editCatalog/", ui.EditCatalogHandler, auth.AdminAuthVerify()),

		// Link管理
		router.NewRoute(router.GET, "/admin/content/manageLink/", ui.ManageLinkHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryAllLink/", ui.QueryAllLinkHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryLink/", ui.QueryLinkHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/deleteLink/", ui.DeleteLinkHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/ajaxLink/", ui.AjaxLinkHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/editLink/", ui.EditLinkHandler, auth.AdminAuthVerify()),

		// Image管理
		router.NewRoute(router.GET, "/admin/content/manageImage/", ui.ManageImageHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryAllImage/", ui.QueryAllImageHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/queryImage/", ui.QueryImageHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.GET, "/admin/content/deleteImage/", ui.DeleteImageHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/ajaxImage/", ui.AjaxImageHandler, auth.AdminAuthVerify()),
		router.NewRoute(router.POST, "/admin/content/editImage/", ui.EditImageHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动Content模块
func (instance *content) Startup() bool {
	return true
}

// Cleanup 清除Content模块
func (instance *content) Cleanup() {

}

// Invoke 执行外部命令
func (instance *content) Invoke(param interface{}) bool {
	return false
}
