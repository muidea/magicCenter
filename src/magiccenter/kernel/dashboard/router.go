package dashboard

import (
	"magiccenter/kernel/auth"
	"magiccenter/kernel/dashboard/ui"
	"magiccenter/router"
)

// RegisterRouter 注册路由
func RegisterRouter() {
	router.AddGetRoute("/admin/", ui.AdminViewHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/admin/login/", ui.LoginViewHandler, nil)
	router.AddPostRoute("/admin/verify/", ui.VerifyAuthActionHandler, nil)
	router.AddGetRoute("/admin/logout/", ui.LogoutActionHandler, auth.AdminAuthVerify())

	//module.RegisterRouter()

	//system.RegisterRouter()
}
