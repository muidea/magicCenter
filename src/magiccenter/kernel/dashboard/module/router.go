package module

import (
	"magiccenter/kernel/auth"
	"magiccenter/kernel/dashboard/module/ui"
	"magiccenter/router"
)

func RegisterRouter() {
	// module管理主视图
	router.AddGetRoute("/admin/system/moduleManage/", ui.ModuleManageHandler, auth.AdminAuthVerify())
	// 启用停用Module
	router.AddPostRoute("/admin/system/applyModuleSetting/", ui.ApplyModuleSettingHandler, auth.AdminAuthVerify())

	// 查询当个Module详细信息
	router.AddGetRoute("/admin/system/queryModuleDetail/", ui.QueryModuleDetailHandler, auth.AdminAuthVerify())
	// 删除Module的Block
	router.AddGetRoute("/admin/system/deleteModuleBlock/", ui.DeleteModuleBlockHandler, auth.AdminAuthVerify())
	// 保存Module的Block
	router.AddPostRoute("/admin/system/ajaxModuleBlock/", ui.SaveModuleBlockHandler, auth.AdminAuthVerify())

	// Page管理主视图
	router.AddGetRoute("/admin/system/modulePage/", ui.ModulePageHandler, auth.AdminAuthVerify())
	// 保存Page的Block
	router.AddPostRoute("/admin/system/ajaxPageBlock/", ui.SavePageBlockHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/admin/system/moduleContent/", ui.ModuleContentHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/system/queryModuleContent/", ui.QueryModuleContentHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/system/ajaxBlockItem/", ui.SaveBlockContentHandler, auth.AdminAuthVerify())
}
