package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// AuthorityHandler 鉴权处理器
type AuthorityHandler interface {
	// 校验权限
	VerifyAuthority(res http.ResponseWriter, req *http.Request) bool

	// 查询指定Route的授权组信息
	QueryRouteAuthGroup(id, pattern string) []model.AuthGroup
	// 更新指定Route的授权组信息
	UpdateRouteAuthGroup(id, pattern string, authGroups []model.AuthGroup)

	// 查询指定用户的授权组信息
	QueryUserAuthGroup(user model.User) []int
	// 更新指定用户的授权信息
	UpdateUserAuthGroup(user model.User, authGroups []int)

	// 查询指定用户使用的模块信息
	QueryUserModule(user model.User) []string
	// 更新指定用户使用的模块信息
	UpdateUserModule(User model.User, modules []string)
}
