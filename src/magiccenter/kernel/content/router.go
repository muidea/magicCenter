package content

import (
	"magiccenter/router"
	"magiccenter/kernel/content/ui"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/content/manageArticle/", ui.ManageArticleHandler)
	router.AddGetRoute("/admin/content/queryAllArticle/", ui.QueryAllArticleHandler)
	router.AddGetRoute("/admin/content/queryArticle/", ui.QueryArticleHandler)
	router.AddGetRoute("/admin/content/deleteArticle/", ui.DeleteArticleHandler)
	router.AddPostRoute("/admin/content/ajaxArticle/", ui.AjaxArticleHandler)	
	router.AddGetRoute("/admin/content/editArticle/", ui.EditArticleHandler)
	
	router.AddGetRoute("/admin/content/manageCatalog/", ui.ManageCatalogHandler)
	router.AddGetRoute("/admin/content/queryAllCatalog/", ui.QueryAllCatalogHandler)
	router.AddGetRoute("/admin/content/queryCatalog/", ui.QueryCatalogHandler)
	router.AddGetRoute("/admin/content/deleteCatalog/", ui.DeleteCatalogHandler)
	router.AddPostRoute("/admin/content/ajaxCatalog/", ui.AjaxCatalogHandler)
	router.AddGetRoute("/admin/content/editCatalog/", ui.EditCatalogHandler)

	router.AddGetRoute("/admin/content/manageLink/", ui.ManageLinkHandler)
	router.AddGetRoute("/admin/content/queryAllLink/", ui.QueryAllLinkHandler)
	router.AddGetRoute("/admin/content/queryLink/", ui.QueryLinkHandler)
	router.AddGetRoute("/admin/content/deleteLink/", ui.DeleteLinkHandler)
	router.AddPostRoute("/admin/content/ajaxLink/", ui.AjaxLinkHandler)	
	router.AddGetRoute("/admin/content/editLink/", ui.EditLinkHandler)
	
	router.AddGetRoute("/admin/content/manageImage/", ui.ManageImageHandler)
	router.AddGetRoute("/admin/content/queryAllImage/", ui.QueryAllImageHandler)
	router.AddGetRoute("/admin/content/deleteImage/", ui.DeleteImageHandler)
	router.AddPostRoute("/admin/content/ajaxImage/", ui.AjaxImageHandler)
	router.AddGetRoute("/admin/content/editImage/", ui.EditImageHandler)
}
