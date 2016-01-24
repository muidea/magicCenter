package admin

import (
	"webcenter/kernel"
	"webcenter/kernel/admin/auth"
	"webcenter/kernel/admin/content"
	"webcenter/kernel/admin/system"
)

func init() {
	auth.RegisterRouter()
	
	content.RegisterRouter()

	registerRouter()
}

func registerRouter() {
	kernel.RegisterGetHandler("/admin/", AdminHandler)
	
	kernel.RegisterGetHandler("/admin/system/manageSystem/", system.ManageSystemHandler)
	kernel.RegisterPostHandler("/admin/system/updateSystem/", system.UpdateSystemHandler)
	
	kernel.RegisterGetHandler("/admin/system/manageModule/", system.ManageModuleHandler)
	kernel.RegisterPostHandler("/admin/system/applyModule/", system.ApplyModuleHandler)	
	kernel.RegisterGetHandler("/admin/system/maintainModule/", system.MaintainModuleHandler)
}

