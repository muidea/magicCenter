package def

import "muidea.com/magicCenter/application/common"

// ID 模块ID
const ID = common.AuthorityModuleID

// Name 模块名称
const Name = "Magic Authority"

// Description 模块描述信息
const Description = "Magic 权限管理模块"

// URL 模块Url
const URL = "/authority"

// GetModuleACL 查询指定Module的ACL
const GetModuleACL = "/module/acl/"

// GetACL 查询指定的ACL
const GetACL = "/acl/"

// PostACL 新增ACL
const PostACL = "/acl/"

// DeleteACL 删除ACL
const DeleteACL = "/acl/:id"

// PutACL 更新ACL（启用、禁用）
const PutACL = "/acl/"

// GetACLAuthGroup 查询指定acl的权限组
const GetACLAuthGroup = "/acl/authgroup/"

// PutACLAuthGroup 更新指定acl的权限组
const PutACLAuthGroup = "/acl/authgroup/"

// GetUserModule 查询指定用户拥有的Module
const GetUserModule = "/user/module/"

// PutUserModule 更新指定用户拥有的Module
const PutUserModule = "/user/module/"

// GetModuleUser 查询拥有指定Module的用户
const GetModuleUser = "/module/user/"

// PutModuleUser 更新拥有指定Module的用户
const PutModuleUser = "/module/user/"
