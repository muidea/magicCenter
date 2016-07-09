package module

import (
	"magiccenter/kernel/auth"
	"magiccenter/kernel/module/ui"
	"magiccenter/router"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/system/moduleManage/", ui.ModuleManageHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/system/applyModuleSetting/", ui.ApplyModuleSettingHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/system/queryModuleDetail/", ui.QueryModuleDetailHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/system/deleteModuleBlock/", ui.DeleteModuleBlockHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/system/ajaxModuleBlock/", ui.SaveModuleBlockHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/system/ajaxPageBlock/", ui.SavePageBlockHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/admin/system/modulePage/", ui.ModulePageHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/admin/system/moduleContent/", ui.ModuleContentHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/system/queryModuleContent/", ui.QueryModuleContentHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/system/ajaxBlockItem/", ui.SaveBlockContentHandler, auth.AdminAuthVerify())
}
