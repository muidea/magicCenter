package system


import (
	"magiccenter/router"
	"magiccenter/kernel/system/ui"
	"magiccenter/kernel/auth"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/system/manageSystem/", ui.ManageSystemHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/system/updateSystem/", ui.UpdateSystemHandler, auth.AdminAuthVerify())
}
