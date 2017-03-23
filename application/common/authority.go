package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// AuthTokenID 鉴权Token
const AuthTokenID = "authToken"

// AuthorityHandler 鉴权处理Handler
type AuthorityHandler interface {
	//@in account 账号
	//@in password 密码
	//@ret model.UserDetail 登陆用户
	//@ret string 本次登陆的鉴权token
	//@ret bool 是否登陆成功
	LoginAccount(account, password string) (model.UserDetail, string, bool)

	//@in authToken 鉴权token
	//@ret bool 是否登出成功
	LogoutAccount(authToken string) bool

	// 校验权限是否OK
	VerifyAuth(res http.ResponseWriter, req *http.Request) bool

	// 调整用户授权组
	AdjustUserAuthGroup(userID int, authGroup []int) bool
	// 获取指定用户的授权组
	GetUserAuthGroup(userID int) ([]int, bool)

	// 更新acl route
	AddACLRoute(moduleURL string, route Route) bool
	DelACLRoute(moduleURL string, route Route) bool

	// 调整acl的授权组
	AdjustACLAuthGroup(modelURL string, route Route, authGroup []int) bool
}
