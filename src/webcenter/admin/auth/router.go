package auth

import (
	"webcenter/application"
	"webcenter/admin/auth/account"
	"webcenter/admin/auth/group"
)

func init() {
	registerRouter()
}

func registerRouter() {
	application.RegisterGetHandler("/auth/login/", LoginHandler)
	application.RegisterGetHandler("/auth/logout/", LogoutHandler)
    application.RegisterPostHandler("/auth/verify/", VerifyHandler)    
	
    application.RegisterGetHandler("/admin/account/manageUser/", account.ManageUserHandler)
	application.RegisterPostHandler("/admin/account/queryAllUser/", account.QueryAllUserHandler)
	application.RegisterPostHandler("/admin/account/checkAccount/", account.CheckAccountHandler)
	application.RegisterPostHandler("/admin/account/queryUser/", account.QueryUserHandler)
	application.RegisterPostHandler("/admin/account/deleteUser/", account.DeleteUserHandler)
    application.RegisterPostHandler("/admin/account/ajaxUser/", account.AjaxUserHandler)
    application.RegisterPostHandler("/admin/account/editUser/", account.EditUserHandler)
    application.RegisterGetHandler("/user/verify/", account.VerifyAccountHandler)
    application.RegisterPostHandler("/admin/account/ajaxVerifyUser/", account.AjaxVerifyUserHandler)
    
    application.RegisterGetHandler("/admin/account/manageGroup/", group.ManageGroupHandler)
	application.RegisterPostHandler("/admin/account/queryAllGroup/", group.QueryAllGroupHandler)
	application.RegisterPostHandler("/admin/account/queryGroup/", group.QueryGroupHandler)
	application.RegisterPostHandler("/admin/account/deleteGroup/", group.DeleteGroupHandler)    
    application.RegisterPostHandler("/admin/account/ajaxGroup/", group.AjaxGroupHandler)
    application.RegisterPostHandler("/admin/account/editGroup/", group.EditGroupHandler)
}