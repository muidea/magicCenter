package auth

import (
	"webcenter/application"
	"webcenter/auth/account"
	"webcenter/auth/group"
	"webcenter/auth/access"
)

func init() {
	registerRouter()
}

func registerRouter() {
	application.RegisterGetHandler("/auth/login/", access.LoginHandler)
	application.RegisterGetHandler("/auth/logout/", access.LogoutHandler)
    application.RegisterPostHandler("/auth/verify/", access.VerifyHandler)
    
	application.RegisterGetHandler("/admin/account/queryAllUser/", account.QueryAllUserHandler)
	application.RegisterGetHandler("/admin/account/queryUser/", account.QueryUserHandler)
	application.RegisterPostHandler("/admin/account/deleteUser/", account.DeleteUserHandler)
    application.RegisterPostHandler("/admin/account/ajaxUser/", account.AjaxUserHandler)
    
	application.RegisterGetHandler("/admin/account/queryAllGroup/", group.QueryAllGroupHandler)
	application.RegisterGetHandler("/admin/account/queryGroup/", group.QueryGroupHandler)
	application.RegisterPostHandler("/admin/account/deleteGroup/", group.DeleteGroupHandler)    
    application.RegisterPostHandler("/admin/account/ajaxGroup/", group.AjaxGroupHandler)
}
