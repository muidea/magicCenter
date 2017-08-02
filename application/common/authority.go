package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// AuthorityHandler 鉴权处理器
type AuthorityHandler interface {
	VerifyAuthority(res http.ResponseWriter, req *http.Request) bool

	// 更新授权组
	QueryAuthGroup(module string) ([]model.AuthGroup, bool)
	InsertAuthGroup(authGroups []model.AuthGroup) bool
	DeleteAuthGroup(authGroups []model.AuthGroup) bool

	// 调整用户授权组
	AdjustUserAuthGroup(userID int, authGroup []int) bool
	// 获取指定用户的授权组
	GetUserAuthGroup(userID int) ([]int, bool)

	// 查询Acl， module=all表示查询所有, status 表示状态0未激活，1激活
	QueryACL(module string, status int) ([]model.ACL, bool)
	// 更新acl
	AddACL(url, method, module string) (model.ACL, bool)
	DelACL(url, method, module string) bool
	EnableACL(acls []int) bool
	DisableACL(acls []int) bool

	// 调整acl的授权组
	AdjustACLAuthGroup(aid int, authGroup []int) (model.ACL, bool)
}
