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

// QueryACL 查询ACL
const QueryACL = "/acl/"

// GetACLByID 查询指定的ACL
const GetACLByID = "/acl/:id"

// PostACL 新增ACL
const PostACL = "/acl/"

// DeleteACL 删除ACL
const DeleteACL = "/acl/:id"

// PutACL 更新ACL
const PutACL = "/acl/:id"

// PutACLs 批量更新ACL（启用、禁用）
const PutACLs = "/acls/"

// GetACLAuthGroup 查询指定acl的权限组
const GetACLAuthGroup = "/acl/authgroup/:id"

// PutACLAuthGroup 更新指定acl的权限组
const PutACLAuthGroup = "/acl/authgroup/:id"

// QueryModule 查询Module的用户信息
const QueryModule = "/module/"

// GetModuleByID 查询指定Module的用户授权组信息
const GetModuleByID = "/module/:id"

// PutModule 更新指定Module的用户授权组信息
const PutModule = "/module/:id"

// QueryUser 查询用户的Module信息
const QueryUser = "/user/"

// GetUserByID 查询指定User的Module授权组信息
const GetUserByID = "/user/:id"

// PutUser 更新指定用户的Module授权组信息
const PutUser = "/user/:id"
