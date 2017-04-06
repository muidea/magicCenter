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

	//@in authToken 鉴权token
	//@ret bool 是否登出成功
	LogoutAccount(authToken string) bool

	// 校验权限是否OK
	VerifyAuth(res http.ResponseWriter, req *http.Request) bool

	// 更新授权组
	QueryAuthGroup(module string) ([]model.AuthGroup, bool)
	InsertAuthGroup(authGroups []model.AuthGroup) bool
	DeleteAuthGroup(authGroups []model.AuthGroup) bool

	// 调整用户授权组
	AdjustUserAuthGroup(userID int, authGroup []int) bool
	// 获取指定用户的授权组
	GetUserAuthGroup(userID int) ([]int, bool)

	// 查询Acl， module=all表示查询所有
	QueryACL(module string) ([]model.ACL, bool)
	// 更新acl
	AddACL(url, method, module string) (model.ACL, bool)
	DelACL(url, method, module string) bool

	// 调整acl的授权组
	AdjustACLAuthGroup(url, method, module string, authGroup []int) (model.ACL, bool)
}
