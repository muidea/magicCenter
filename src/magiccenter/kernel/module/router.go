package module

import (
	"magiccenter/router"
	"magiccenter/kernel/module/ui"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/system/moduleManage/", ui.ModuleManageHandler)
	router.AddPostRoute("/admin/system/applyModuleSetting/", ui.ApplyModuleSettingHandler)
	router.AddGetRoute("/admin/system/queryModuleDetail/", ui.QueryModuleDetailHandler)
	router.AddGetRoute("/admin/system/deleteModuleBlock/", ui.DeleteModuleBlockHandler)
	router.AddPostRoute("/admin/system/ajaxModuleBlock/", ui.SaveModuleBlockHandler)	
	router.AddPostRoute("/admin/system/ajaxPageBlock/", ui.SavePageBlockHandler)
		
	router.AddGetRoute("/admin/system/moduleContent/", ui.ModuleContentHandler)
	router.AddGetRoute("/admin/system/queryModuleContent/", ui.QueryModuleContentHandler)
}

