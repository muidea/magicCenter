package account

import (
	"magiccenter/kernel/account/ui"
	"magiccenter/kernel/auth"
	"magiccenter/router"
)

func RegisterRouter() {
	// 用户管理主视图页面
	router.AddGetRoute("/admin/account/manageUser/", ui.ManageUserHandler, auth.AdminAuthVerify())
	// 查询全部用户信息
	router.AddGetRoute("/admin/account/queryAllUser/", ui.QueryAllUserHandler, auth.AdminAuthVerify())
	// 查询指定用户信息
	router.AddGetRoute("/admin/account/queryUser/", ui.QueryUserHandler, auth.AdminAuthVerify())
	// 检查Account是否可用
	router.AddPostRoute("/admin/account/checkAccount/", ui.CheckAccountHandler, auth.AdminAuthVerify())
	// 删除指定用户
	router.AddGetRoute("/admin/account/deleteUser/", ui.DeleteUserHandler, auth.AdminAuthVerify())
	// 更新用户信息（新建用户或者更新指定用户的信息）
	router.AddPostRoute("/admin/account/ajaxUser/", ui.AjaxUserHandler, auth.AdminAuthVerify())

	// 用户组管理主视图页面
	router.AddGetRoute("/admin/account/manageGroup/", ui.ManageGroupHandler, auth.AdminAuthVerify())
	// 查询全部用户组信息
	router.AddGetRoute("/admin/account/queryAllGroup/", ui.QueryAllGroupHandler, auth.AdminAuthVerify())
	// 查询指定用户组
	router.AddGetRoute("/admin/account/queryGroup/", ui.QueryGroupHandler, auth.AdminAuthVerify())
	// 删除指定用户组
	router.AddGetRoute("/admin/account/deleteGroup/", ui.DeleteGroupHandler, auth.AdminAuthVerify())
	// 更新用户组信息（新建用户组活更新指定用户组）
	router.AddPostRoute("/admin/account/ajaxGroup/", ui.AjaxGroupHandler, auth.AdminAuthVerify())

	// 用户空间视图
	router.AddGetRoute("/user/profile/", ui.UserProfileViewHandler, nil)
	// 校验用户信息视图，用于用户补充昵称，密码这些额外信息
	router.AddGetRoute("/user/verify/", ui.UserVerifyViewHandler, nil)
	// 更新用户信息
	router.AddPostRoute("/user/ajaxVerify/", ui.AjaxVerifyHandler, nil)
}
