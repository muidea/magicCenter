package def

import "muidea.com/magicCenter/common"

// ID 模块ID
const ID = common.EndpointModuleID

// Name 模块名称
const Name = "Magic Endpoint"

// Description 模块描述信息
const Description = "Magic 终端管理模块"

// URL 模块Url
const URL = "/endpoint"

// QueryEndpoint 查询Endpoint
const QueryEndpoint = "/registry/"

// QueryByIDEndpoint 查询Endpoint
const QueryByIDEndpoint = "/registry/:id"

// PostEndpoint 新增Endpoint
const PostEndpoint = "/registry/"

// PutEndpoint 更新指定Endpoint
const PutEndpoint = "/registry/:id"

// DeleteEndpoint 删除指定Endpoint
const DeleteEndpoint = "/registry/:id"
