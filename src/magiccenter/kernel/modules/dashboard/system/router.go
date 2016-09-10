package system

import (
	"magiccenter/kernel/auth"
	"magiccenter/kernel/dashboard/system/ui"
	"magiccenter/router"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/system/manageSystem/", ui.ManageSystemHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/system/updateSystem/", ui.UpdateSystemHandler, auth.AdminAuthVerify())
}
