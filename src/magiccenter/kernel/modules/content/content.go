package content

import (
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	commonmodel "magiccenter/common/model"
	"magiccenter/kernel/modules/content/bll"
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
const URL string = "/content"

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

func (instance *content) Status() int {
	return 0
}

func (instance *content) EndPoint() common.EndPoint {
	return nil
}

func (instance *content) AuthGroups() []common.AuthGroup {
	groups := []common.AuthGroup{}

	return groups
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
	switch param.(type) {
	case *commonbll.QueryContentMetaRequest:
		{
			request := param.(*commonbll.QueryContentMetaRequest)
			if request != nil {
				response := result.(*commonbll.QueryContentMetaResponse)
				articleMeta := commonmodel.ContentMeta{Subject: commonmodel.ARTICLE, Description: "文章", URL: "conent/article/"}
				response.ContentMetas = append(response.ContentMetas, articleMeta)

				catalogMeta := commonmodel.ContentMeta{Subject: commonmodel.CATALOG, Description: "分类", URL: "conent/catalog/"}
				response.ContentMetas = append(response.ContentMetas, catalogMeta)

				linkMeta := commonmodel.ContentMeta{Subject: commonmodel.LINK, Description: "链接", URL: "conent/link/"}
				response.ContentMetas = append(response.ContentMetas, linkMeta)

				mediaMeta := commonmodel.ContentMeta{Subject: commonmodel.MEDIA, Description: "文件", URL: "conent/media/"}
				response.ContentMetas = append(response.ContentMetas, mediaMeta)
				return true
			}

			return false
		}
	case *commonbll.QueryContentArticleRequest:
		{
			request := param.(*commonbll.QueryContentArticleRequest)
			if request != nil {
				response := result.(*commonbll.QueryContentArticleResponse)
				response.Articles = bll.QueryAllArticleSummary()
				return true
			}

			return false
		}
	case *commonbll.QuerySingleArticleRequest:
		{
			request := param.(*commonbll.QuerySingleArticleRequest)
			if request != nil {
				response := result.(*commonbll.QuerySingleArticleResponse)
				response.Article, response.Found = bll.QueryArticleByID(request.ID)
				return true
			}

			return false
		}
	case *commonbll.CreateArticleRequest:
		{
			request := param.(*commonbll.CreateArticleRequest)
			if request != nil {
				response := result.(*commonbll.CreateArticleResponse)
				response.Article, response.Result = bll.CreateArticle(request.Title, request.Content, request.Creater, request.Catalog)
				return true
			}

			return false
		}
	case *commonbll.UpdateArticleRequest:
		{
			request := param.(*commonbll.UpdateArticleRequest)
			if request != nil {
				response := result.(*commonbll.UpdateArticleResponse)
				response.Article, response.Result = bll.SaveArticle(request.Article)
				return true
			}

			return false
		}
	case *commonbll.DeleteArticleRequest:
		{
			request := param.(*commonbll.DeleteArticleRequest)
			if request != nil {
				response := result.(*commonbll.DeleteArticleResponse)
				response.Result = bll.DeleteArticle(request.ID)
				return true
			}

			return false
		}
	case *commonbll.QueryContentCatalogRequest:
		{
			request := param.(*commonbll.QueryContentCatalogRequest)
			if request != nil {
				response := result.(*commonbll.QueryContentCatalogResponse)
				response.Catalogs = bll.QueryAllCatalog()
				return true
			}

			return false
		}
	case *commonbll.QueryContentLinkRequest:
		{
			request := param.(*commonbll.QueryContentLinkRequest)
			if request != nil {
				response := result.(*commonbll.QueryContentLinkResponse)
				response.Links = bll.QueryAllLink()
				return true
			}

			return false
		}
	case *commonbll.QueryContentMediaRequest:
		{
			request := param.(*commonbll.QueryContentMediaRequest)
			if request != nil {
				response := result.(*commonbll.QueryContentMediaResponse)
				response.Medias = bll.QueryAllMedia()
				return true
			}

			return false
		}
	}

	return false
}
