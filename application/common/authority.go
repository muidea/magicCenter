package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// AuthorityHandler 鉴权处理器
type AuthorityHandler interface {
	// 校验权限
	VerifyAuthority(res http.ResponseWriter, req *http.Request) bool

	// 查询指定Module的ACL
	QueryModuleACL(module string) []model.ACL
	// 查询指定ACL
	QueryACL(url, method string) (model.ACL, bool)
	// 新增ACL
	InsertACL(url, method, module string, status int) (model.ACL, bool)
	// 删除ACL
	DeleteACL(id int) bool
	// 启用ACL
	EnableACL(ids []int) bool
	// 禁用ACL
	DisableACL(ids []int) bool

	// 查询指定ACL的授权组信息
	QueryACLAuthGroup(id int) []int
	// 更新指定ACL的授权组信息
	UpdateACLAuthGroup(id int, authGroups []int) bool

	// 查询指定用户的授权组信息
	QueryUserAuthGroup(user int) []int
	// 更新指定用户的授权信息
	UpdateUserAuthGroup(user int, authGroups []int) bool

	// 查询指定用户使用的模块信息
	QueryUserModule(user int) []string
	// 更新指定用户使用的模块信息
	UpdateUserModule(user int, modules []string) bool
}
