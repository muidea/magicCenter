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
		// Article管理视图
		router.NewRoute(common.GET, "manageArticleView/", ui.ManageArticleViewHandler, auth.AdminAuthVerify()),
		// Catalog管理视图
		router.NewRoute(common.GET, "manageCatalogView/", ui.ManageCatalogViewHandler, auth.AdminAuthVerify()),
		// Link管理视图
		router.NewRoute(common.GET, "manageLinkView/", ui.ManageLinkViewHandler, auth.AdminAuthVerify()),
		// Media管理视图
		router.NewRoute(common.GET, "manageMediaView/", ui.ManageMediaViewHandler, auth.AdminAuthVerify()),
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
	case *commonbll.QueryContentArticleListRequest:
		{
			request := param.(*commonbll.QueryContentArticleListRequest)
			if request != nil {
				response := result.(*commonbll.QueryContentArticleListResponse)
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
	case *commonbll.QueryContentCatalogListRequest:
		{
			request := param.(*commonbll.QueryContentCatalogListRequest)
			if request != nil {
				response := result.(*commonbll.QueryContentCatalogListResponse)
				response.Catalogs = bll.QueryAllCatalog()
				return true
			}

			return false
		}
	case *commonbll.QuerySingleCatalogRequest:
		{
			request := param.(*commonbll.QuerySingleCatalogRequest)
			if request != nil {
				response := result.(*commonbll.QuerySingleCatalogResponse)
				response.Catalog, response.Found = bll.QueryCatalogByID(request.ID)
				return true
			}

			return false
		}
	case *commonbll.CreateCatalogRequest:
		{
			request := param.(*commonbll.CreateCatalogRequest)
			if request != nil {
				response := result.(*commonbll.CreateCatalogResponse)
				response.Catalog, response.Result = bll.CreateCatalog(request.Name, request.Creater, request.Parent)
				return true
			}

			return false
		}
	case *commonbll.UpdateCatalogRequest:
		{
			request := param.(*commonbll.UpdateCatalogRequest)
			if request != nil {
				response := result.(*commonbll.UpdateCatalogResponse)
				response.Catalog, response.Result = bll.UpdateCatalog(request.Catalog)
				return true
			}

			return false
		}
	case *commonbll.DeleteCatalogRequest:
		{
			request := param.(*commonbll.DeleteCatalogRequest)
			if request != nil {
				response := result.(*commonbll.DeleteCatalogResponse)
				response.Result = bll.DeleteCatalog(request.ID)
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
	case *commonbll.QuerySingleLinkRequest:
		{
			request := param.(*commonbll.QuerySingleLinkRequest)
			if request != nil {
				response := result.(*commonbll.QuerySingleLinkResponse)
				response.Link, response.Found = bll.QueryLinkByID(request.ID)
				return true
			}

			return false
		}
	case *commonbll.CreateLinkRequest:
		{
			request := param.(*commonbll.CreateLinkRequest)
			if request != nil {
				response := result.(*commonbll.CreateLinkResponse)
				response.Link, response.Result = bll.CreateLink(request.Name, request.URL, request.Logo, request.Creater, request.Catalog)
				return true
			}

			return false
		}
	case *commonbll.UpdateLinkRequest:
		{
			request := param.(*commonbll.UpdateLinkRequest)
			if request != nil {
				response := result.(*commonbll.UpdateLinkResponse)
				response.Link, response.Result = bll.SaveLink(request.Link)
				return true
			}

			return false
		}
	case *commonbll.DeleteLinkRequest:
		{
			request := param.(*commonbll.DeleteLinkRequest)
			if request != nil {
				response := result.(*commonbll.DeleteLinkResponse)
				response.Result = bll.DeleteLinkByID(request.ID)
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
	case *commonbll.QuerySingleMediaRequest:
		{
			request := param.(*commonbll.QuerySingleMediaRequest)
			if request != nil {
				response := result.(*commonbll.QuerySingleMediaResponse)
				response.Media, response.Found = bll.QueryMediaByID(request.ID)
				return true
			}

			return false
		}
	case *commonbll.CreateMediaRequest:
		{
			request := param.(*commonbll.CreateMediaRequest)
			if request != nil {
				response := result.(*commonbll.CreateMediaResponse)
				response.Media, response.Result = bll.CreateMedia(request.Name, request.URL, request.Type, request.Desc, request.Creater, request.Catalog)
				return true
			}

			return false
		}
	case *commonbll.UpdateMediaRequest:
		{
			request := param.(*commonbll.UpdateMediaRequest)
			if request != nil {
				response := result.(*commonbll.UpdateMediaResponse)
				response.Media, response.Result = bll.SaveMedia(request.Media)
				return true
			}

			return false
		}
	case *commonbll.DeleteMediaRequest:
		{
			request := param.(*commonbll.DeleteMediaRequest)
			if request != nil {
				response := result.(*commonbll.DeleteMediaResponse)
				response.Result = bll.DeleteMediaByID(request.ID)
				return true
			}

			return false
		}
	}

	return false
}