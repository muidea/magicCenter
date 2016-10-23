package content

import (
	"magiccenter/common"
	"magiccenter/kernel/modules/content/ui"
	"magiccenter/system"

	"muidea.com/util"
)

// ID 模块ID
const ID = "3a7123ec-63f0-5e46-1234-e6ca1af6fe4e"

// Name 模块名称
const Name = "Magic Content"

// Description 模块描述信息
const Description = "Magic 内容管理模块"

// URL 模块Url
const URL string = "content"

type content struct {
}

var instance *content

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &content{}
	}

	modulehub := system.GetModuleHub()
	modulehub.RegisterModule(instance)
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

func (instance *content) EndPoint() common.EndPoint {
	return nil
}

// Route 路由信息
func (instance *content) Routes() []common.Route {
	router := system.GetRouter()
	auth := system.GetAuthority()

	routes := []common.Route{
		// 管理Article视图
		router.NewRoute(common.GET, "manageArticleView/", ui.ManageArticleViewHandler, auth.AdminAuthVerify()),
		// 查询全部ArticleSummary
		router.NewRoute(common.GET, "queryAllArticleSummary/", ui.QueryAllArticleSummaryHandler, auth.AdminAuthVerify()),
		// 查询指定Article
		router.NewRoute(common.GET, "queryArticle/", ui.QueryArticleHandler, auth.AdminAuthVerify()),
		// 删除指定Article
		router.NewRoute(common.GET, "deleteArticle/", ui.DeleteArticleHandler, auth.AdminAuthVerify()),
		// 保存文章
		router.NewRoute(common.POST, "ajaxArticle/", ui.AjaxArticleHandler, auth.AdminAuthVerify()),

		// Catalog管理视图
		router.NewRoute(common.GET, "manageCatalogView/", ui.ManageCatalogViewHandler, auth.AdminAuthVerify()),
		// 查询全部Catalog
		router.NewRoute(common.GET, "queryAllCatalog/", ui.QueryAllCatalogHandler, auth.AdminAuthVerify()),
		// 查询指定Catalog
		router.NewRoute(common.GET, "queryCatalog/", ui.QueryCatalogHandler, auth.AdminAuthVerify()),
		// 删除指定Catalog
		router.NewRoute(common.GET, "deleteCatalog/", ui.DeleteCatalogHandler, auth.AdminAuthVerify()),
		// 保存Catalog
		router.NewRoute(common.POST, "ajaxCatalog/", ui.AjaxCatalogHandler, auth.AdminAuthVerify()),

		// Link管理视图
		router.NewRoute(common.GET, "manageLinkView/", ui.ManageLinkViewHandler, auth.AdminAuthVerify()),
		// 查询全部Link
		router.NewRoute(common.GET, "queryAllLink/", ui.QueryAllLinkHandler, auth.AdminAuthVerify()),
		// 查询指定Link
		router.NewRoute(common.GET, "queryLink/", ui.QueryLinkHandler, auth.AdminAuthVerify()),
		// 删除指定Link
		router.NewRoute(common.GET, "deleteLink/", ui.DeleteLinkHandler, auth.AdminAuthVerify()),
		// 保存Link
		router.NewRoute(common.POST, "ajaxLink/", ui.AjaxLinkHandler, auth.AdminAuthVerify()),

		// Media管理视图
		router.NewRoute(common.GET, "manageMediaView/", ui.ManageMediaViewHandler, auth.AdminAuthVerify()),
		// 查询全部Media
		router.NewRoute(common.GET, "queryAllMedia/", ui.QueryAllMediaHandler, auth.AdminAuthVerify()),
		// 查询指定Media
		router.NewRoute(common.GET, "queryMedia/", ui.QueryMediaHandler, auth.AdminAuthVerify()),
		// 删除指定Media
		router.NewRoute(common.GET, "deleteMedia/", ui.DeleteMediaHandler, auth.AdminAuthVerify()),
		// 保存Media
		router.NewRoute(common.POST, "ajaxMedia/", ui.AjaxMediaHandler, auth.AdminAuthVerify()),
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
func (instance *content) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}
