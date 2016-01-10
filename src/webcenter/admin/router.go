package admin

import (
	"webcenter/application"
	"webcenter/admin/system"
)

func init() {
	registerRouter()
}

func registerRouter() {
	application.RegisterGetHandler("/admin/", AdminHandler)
	
	application.RegisterGetHandler("/admin/system/manageSystem/", system.ManageSystemHandler)
	application.RegisterPostHandler("/admin/system/updateSystem/", system.UpdateSystemHandler)
	
	application.RegisterGetHandler("/admin/system/manageModule/", system.ManageModuleHandler)
	application.RegisterPostHandler("/admin/system/updateModule/", system.UpdateModuleHandler)	
}

