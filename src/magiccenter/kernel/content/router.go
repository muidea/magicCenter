package content

import (
	"magiccenter/router"
	"magiccenter/kernel/content/ui"
	"magiccenter/kernel/auth"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/content/manageArticle/", ui.ManageArticleHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/queryAllArticle/", ui.QueryAllArticleHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/queryArticle/", ui.QueryArticleHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/deleteArticle/", ui.DeleteArticleHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/content/ajaxArticle/", ui.AjaxArticleHandler, auth.AdminAuthVerify())	
	router.AddGetRoute("/admin/content/editArticle/", ui.EditArticleHandler, auth.AdminAuthVerify())
	
	router.AddGetRoute("/admin/content/manageCatalog/", ui.ManageCatalogHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/queryAllCatalog/", ui.QueryAllCatalogHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/queryCatalog/", ui.QueryCatalogHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/deleteCatalog/", ui.DeleteCatalogHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/content/ajaxCatalog/", ui.AjaxCatalogHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/editCatalog/", ui.EditCatalogHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/admin/content/manageLink/", ui.ManageLinkHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/queryAllLink/", ui.QueryAllLinkHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/queryLink/", ui.QueryLinkHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/deleteLink/", ui.DeleteLinkHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/content/ajaxLink/", ui.AjaxLinkHandler, auth.AdminAuthVerify())	
	router.AddGetRoute("/admin/content/editLink/", ui.EditLinkHandler, auth.AdminAuthVerify())
	
	router.AddGetRoute("/admin/content/manageImage/", ui.ManageImageHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/queryAllImage/", ui.QueryAllImageHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/deleteImage/", ui.DeleteImageHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/content/ajaxImage/", ui.AjaxImageHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/content/editImage/", ui.EditImageHandler, auth.AdminAuthVerify())
}
