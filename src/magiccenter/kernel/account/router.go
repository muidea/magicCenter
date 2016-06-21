package account

import (
	"magiccenter/kernel/account/ui"
	"magiccenter/kernel/auth"
	"magiccenter/router"
)

func RegisterRouter() {
	router.AddGetRoute("/admin/account/manageUser/", ui.ManageUserHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/account/queryAllUser/", ui.QueryAllUserHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/account/queryUser/", ui.QueryUserHandler, auth.AdminAuthVerify())
	// 检查Account是否可用
	router.AddPostRoute("/admin/account/checkAccount/", ui.CheckAccountHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/account/deleteUser/", ui.DeleteUserHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/account/ajaxUser/", ui.AjaxUserHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/admin/account/manageGroup/", ui.ManageGroupHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/account/queryAllGroup/", ui.QueryAllGroupHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/account/queryGroup/", ui.QueryGroupHandler, auth.AdminAuthVerify())
	router.AddGetRoute("/admin/account/deleteGroup/", ui.DeleteGroupHandler, auth.AdminAuthVerify())
	router.AddPostRoute("/admin/account/ajaxGroup/", ui.AjaxGroupHandler, auth.AdminAuthVerify())

	router.AddGetRoute("/user/profile/", ui.UserProfileHandler, auth.LoginAuthVerify())
	router.AddGetRoute("/user/verify/", ui.UserVerifyHandler, nil)
	router.AddPostRoute("/user/ajaxVerify/", ui.AjaxVerifyHandler, nil)
}
