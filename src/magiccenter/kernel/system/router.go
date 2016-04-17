package system


import (
	"magiccenter/router"
	"magiccenter/kernel/system/ui"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/system/manageSystem/", ui.ManageSystemHandler)
	router.AddPostRoute("/admin/system/updateSystem/", ui.UpdateSystemHandler)
}
