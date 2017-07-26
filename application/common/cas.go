package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// AuthTokenID 鉴权Token
const AuthTokenID = "authToken"

// CASHandler 鉴权处理Handler
type CASHandler interface {
	//@in account 账号
	//@in password 密码
	//@ret model.UserDetail 登陆用户
	//@ret string 本次登陆的鉴权token
	//@ret bool 是否登陆成功
	LoginAccount(account, password string) (model.UserDetail, string, bool)

	LoginToken(token string) (string, bool)

	//@in authToken 鉴权token
	//@ret bool 是否登出成功
	Logout(authToken string) bool

	// 校验权限是否OK
	VerifyToken(authToken string) bool

	VerifyAccount(res http.ResponseWriter, req *http.Request) bool
}
