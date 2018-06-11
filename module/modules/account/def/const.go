package def

import "muidea.com/magicCenter/common"

// ID 模块ID
const ID = common.AccountModuleID

// Name 模块名称
const Name = "Magic Account"

// Description 模块描述信息
const Description = "Magic 账号管理模块"

// URL 模块Url
const URL = "/account"

// GetAllUser 查询所有用户
const GetAllUser = "/user/"

// GetUser 查询指定用户
const GetUser = "/user/:id"

// PostUser 新建用户
const PostUser = "/user/"

// PutUser 更新用户信息
const PutUser = "/user/:id"

// DeleteUser 删除用户
const DeleteUser = "/user/:id"

// GetAllGroup 查询所有分组
const GetAllGroup = "/group/"

// GetGroup 查询指定分组
const GetGroup = "/group/:id"

// PostGroup 新建分组
const PostGroup = "/group/"

// PutGroup 更新分组信息
const PutGroup = "/group/:id"

// DeleteGroup 删除分组
const DeleteGroup = "/group/:id"
