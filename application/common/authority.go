package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// AuthorityHandler 鉴权处理器
type AuthorityHandler interface {
	// 校验权限
	VerifyAuthority(res http.ResponseWriter, req *http.Request) bool
	// 查询指定用户的授权组信息
	QueryUserAuthGroup(user model.User) model.AuthGroupInfo
	// 更新指定用户的授权信息
	UpdateUserAuthGroup(user model.User, authGroupInfo model.AuthGroupInfo)
	// 查询指定分组的授权信息
	QueryGroupAuthGroup(group model.Group) model.AuthGroupInfo
	// 更新指定分组的授权信息
	UpdateGroupAuthGroup(group model.Group, authGroupInfo model.AuthGroupInfo)
}
