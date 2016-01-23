package content

import (
	"webcenter/kernel"
	"webcenter/kernel/admin/content/article"
	"webcenter/kernel/admin/content/catalog"
	"webcenter/kernel/admin/content/link"
	"webcenter/kernel/admin/content/image"
)

func RegisterRouter() {
	kernel.RegisterGetHandler("/admin/content/manageArticle/", article.ManageArticleHandler)
	kernel.RegisterPostHandler("/admin/content/queryAllArticle/", article.QueryAllArticleHandler)
	kernel.RegisterPostHandler("/admin/content/queryArticle/", article.QueryArticleHandler)
	kernel.RegisterPostHandler("/admin/content/deleteArticle/", article.DeleteArticleHandler)
	kernel.RegisterPostHandler("/admin/content/ajaxArticle/", article.AjaxArticleHandler)	
	kernel.RegisterPostHandler("/admin/content/editArticle/", article.EditArticleHandler)
	
	kernel.RegisterGetHandler("/admin/content/manageCatalog/", catalog.ManageCatalogHandler)
	kernel.RegisterPostHandler("/admin/content/queryAllCatalogInfo/", catalog.QueryAllCatalogInfoHandler)
	kernel.RegisterPostHandler("/admin/content/queryCatalogInfo/", catalog.QueryCatalogInfoHandler)
	kernel.RegisterPostHandler("/admin/content/queryAvalibleParentCatalogInfo/", catalog.QueryAvalibleParentCatalogInfoHandler)
	kernel.RegisterPostHandler("/admin/content/querySubCatalogInfo/", catalog.QuerySubCatalogInfoHandler)
	kernel.RegisterPostHandler("/admin/content/queryCatalog/", catalog.QueryCatalogHandler)
	kernel.RegisterPostHandler("/admin/content/deleteCatalog/", catalog.DeleteCatalogHandler)
	kernel.RegisterPostHandler("/admin/content/ajaxCatalog/", catalog.AjaxCatalogHandler)	
	kernel.RegisterPostHandler("/admin/content/editCatalog/", catalog.EditCatalogHandler)

	kernel.RegisterGetHandler("/admin/content/manageLink/", link.ManageLinkHandler)
	kernel.RegisterPostHandler("/admin/content/queryAllLink/", link.QueryAllLinkHandler)
	kernel.RegisterPostHandler("/admin/content/queryLink/", link.QueryLinkHandler)
	kernel.RegisterPostHandler("/admin/content/deleteLink/", link.DeleteLinkHandler)
	kernel.RegisterPostHandler("/admin/content/ajaxLink/", link.AjaxLinkHandler)	
	kernel.RegisterPostHandler("/admin/content/editLink/", link.EditLinkHandler)
	
	kernel.RegisterGetHandler("/admin/content/manageImage/", image.ManageImageHandler)
	kernel.RegisterPostHandler("/admin/content/queryAllImage/", image.QueryAllImageHandler)
	kernel.RegisterPostHandler("/admin/content/deleteImage/", image.DeleteImageHandler)
	kernel.RegisterPostHandler("/admin/content/ajaxImage/", image.AjaxImageHandler)
	kernel.RegisterPostHandler("/admin/content/editImage/", image.EditImageHandler)
}
