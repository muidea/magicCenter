package system


import (
	"magiccenter/router"
	"magiccenter/kernel/system/ui"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/system/manageSystem/", ui.ManageSystemHandler)
	router.AddPostRoute("/admin/system/updateSystem/", ui.UpdateSystemHandler)
	
	router.AddGetRoute("/admin/system/manageModule/", ui.ManageModuleHandler)
	router.AddPostRoute("/admin/system/applyModuleSetting/", ui.ApplyModuleSettingHandler)
	router.AddGetRoute("/admin/system/queryModuleDetail/", ui.QueryModuleDetailHandler)
	router.AddGetRoute("/admin/system/deleteModuleBlock/", ui.DeleteModuleBlockHandler)
	router.AddPostRoute("/admin/system/ajaxModuleBlock/", ui.SaveModuleBlockHandler)	
	router.AddPostRoute("/admin/system/ajaxPageBlock/", ui.SavePageBlockHandler)
}
