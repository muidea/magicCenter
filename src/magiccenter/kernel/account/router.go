package account

import (
	"magiccenter/router"
	"magiccenter/kernel/account/ui"
)

func RegisterRouter() {
    router.AddGetRoute("/admin/account/manageUser/", ui.ManageUserHandler)
	router.AddGetRoute("/admin/account/queryAllUser/", ui.QueryAllUserHandler)
	router.AddGetRoute("/admin/account/queryUser/", ui.QueryUserHandler)
	// 检查Account是否可用
	router.AddPostRoute("/admin/account/checkAccount/", ui.CheckAccountHandler)
	router.AddGetRoute("/admin/account/deleteUser/", ui.DeleteUserHandler)
    router.AddPostRoute("/admin/account/ajaxUser/", ui.AjaxUserHandler)
    router.AddPostRoute("/admin/account/updateUser/", ui.UpdateUserHandler)
    
    router.AddGetRoute("/admin/account/manageGroup/", ui.ManageGroupHandler)
	router.AddGetRoute("/admin/account/queryAllGroup/", ui.QueryAllGroupHandler)
	router.AddGetRoute("/admin/account/queryGroup/", ui.QueryGroupHandler)
	router.AddGetRoute("/admin/account/deleteGroup/", ui.DeleteGroupHandler)    
    router.AddPostRoute("/admin/account/ajaxGroup/", ui.AjaxGroupHandler)
    
    router.AddGetRoute("/user/profile/", ui.UserProfileHandler)
    router.AddGetRoute("/user/verify/", ui.UserVerifyHandler)
}
