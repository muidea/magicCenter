package auth

import (
	"webcenter/kernel"
	"webcenter/kernel/admin/auth/account"
	"webcenter/kernel/admin/auth/group"
)

func RegisterRouter() {
	kernel.RegisterGetHandler("/auth/login/", LoginHandler)
	kernel.RegisterGetHandler("/auth/logout/", LogoutHandler)
    kernel.RegisterPostHandler("/auth/verify/", VerifyHandler)    
	
    kernel.RegisterGetHandler("/admin/account/manageUser/", account.ManageUserHandler)
	kernel.RegisterPostHandler("/admin/account/queryAllUser/", account.QueryAllUserHandler)
	kernel.RegisterPostHandler("/admin/account/checkAccount/", account.CheckAccountHandler)
	kernel.RegisterPostHandler("/admin/account/queryUser/", account.QueryUserHandler)
	kernel.RegisterPostHandler("/admin/account/deleteUser/", account.DeleteUserHandler)
    kernel.RegisterPostHandler("/admin/account/ajaxUser/", account.AjaxUserHandler)
    kernel.RegisterPostHandler("/admin/account/editUser/", account.EditUserHandler)
    kernel.RegisterGetHandler("/user/verify/", account.VerifyAccountHandler)
    kernel.RegisterPostHandler("/admin/account/ajaxVerifyUser/", account.AjaxVerifyUserHandler)
    
    kernel.RegisterGetHandler("/admin/account/manageGroup/", group.ManageGroupHandler)
	kernel.RegisterPostHandler("/admin/account/queryAllGroup/", group.QueryAllGroupHandler)
	kernel.RegisterPostHandler("/admin/account/queryGroup/", group.QueryGroupHandler)
	kernel.RegisterPostHandler("/admin/account/deleteGroup/", group.DeleteGroupHandler)    
    kernel.RegisterPostHandler("/admin/account/ajaxGroup/", group.AjaxGroupHandler)
    kernel.RegisterPostHandler("/admin/account/editGroup/", group.EditGroupHandler)
}
