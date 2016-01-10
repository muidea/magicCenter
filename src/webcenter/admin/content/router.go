package content

import (
	"webcenter/application"
	"webcenter/admin/content/article"
	"webcenter/admin/content/catalog"
	"webcenter/admin/content/link"
	"webcenter/admin/content/image"
)

func init() {
	registerRouter()
}

func registerRouter() {
	application.RegisterGetHandler("/admin/content/manageArticle/", article.ManageArticleHandler)
	application.RegisterPostHandler("/admin/content/queryAllArticle/", article.QueryAllArticleHandler)
	application.RegisterPostHandler("/admin/content/queryArticle/", article.QueryArticleHandler)
	application.RegisterPostHandler("/admin/content/deleteArticle/", article.DeleteArticleHandler)
	application.RegisterPostHandler("/admin/content/ajaxArticle/", article.AjaxArticleHandler)	
	application.RegisterPostHandler("/admin/content/editArticle/", article.EditArticleHandler)
	
	application.RegisterGetHandler("/admin/content/manageCatalog/", catalog.ManageCatalogHandler)
	application.RegisterPostHandler("/admin/content/queryAllCatalogInfo/", catalog.QueryAllCatalogInfoHandler)
	application.RegisterPostHandler("/admin/content/queryCatalogInfo/", catalog.QueryCatalogInfoHandler)
	application.RegisterPostHandler("/admin/content/queryAvalibleParentCatalogInfo/", catalog.QueryAvalibleParentCatalogInfoHandler)
	application.RegisterPostHandler("/admin/content/querySubCatalogInfo/", catalog.QuerySubCatalogInfoHandler)
	application.RegisterPostHandler("/admin/content/queryCatalog/", catalog.QueryCatalogHandler)
	application.RegisterPostHandler("/admin/content/deleteCatalog/", catalog.DeleteCatalogHandler)
	application.RegisterPostHandler("/admin/content/ajaxCatalog/", catalog.AjaxCatalogHandler)	
	application.RegisterPostHandler("/admin/content/editCatalog/", catalog.EditCatalogHandler)

	application.RegisterGetHandler("/admin/content/manageLink/", link.ManageLinkHandler)
	application.RegisterPostHandler("/admin/content/queryAllLink/", link.QueryAllLinkHandler)
	application.RegisterPostHandler("/admin/content/queryLink/", link.QueryLinkHandler)
	application.RegisterPostHandler("/admin/content/deleteLink/", link.DeleteLinkHandler)
	application.RegisterPostHandler("/admin/content/ajaxLink/", link.AjaxLinkHandler)	
	application.RegisterPostHandler("/admin/content/editLink/", link.EditLinkHandler)
	
	application.RegisterGetHandler("/admin/content/manageImage/", image.ManageImageHandler)
	application.RegisterPostHandler("/admin/content/queryAllImage/", image.QueryAllImageHandler)
	application.RegisterPostHandler("/admin/content/deleteImage/", image.DeleteImageHandler)
	application.RegisterPostHandler("/admin/content/ajaxImage/", image.AjaxImageHandler)
	application.RegisterPostHandler("/admin/content/editImage/", image.EditImageHandler)
}
