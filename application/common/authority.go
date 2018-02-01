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
	QueryACLByModule(module string) []model.ACL
	// 查询指定ACL
	QueryACLByID(id int) (model.ACL, bool)
	// 新增ACL
	InsertACL(url, method, module string, status int, authGroup int) (model.ACL, bool)

	// 删除ACL
	DeleteACL(id int) bool
	// 更新ACL
	UpdateACLStatus(enableList []int, disableList []int) bool

	// 查询指定ACL的授权组信息
	QueryACLAuthGroup(id int) int
	// 更新指定ACL的授权组信息
	UpdateACLAuthGroup(id int, authGroup int) bool

	// 查询指定模块的拥有者
	QueryModuleUserAuthGroup(module string) model.ModuleUserAuthGroupInfo
	// 更新指定模块的拥有者
	UpdateModuleUserAuthGroup(module string, users []model.UserAuthGroup) bool

	// 查询指定用户的ACL列表
	QueryUserACL(user int) []model.ACL
	// 查询指定用户使用的模块信息
	QueryUserModuleAuthGroup(user int) model.UserModuleAuthGroupInfo
	// 更新指定用户使用的模块信息
	UpdateUserModuleAuthGroup(user int, modules []model.ModuleAuthGroup) bool
}
