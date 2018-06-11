package def

import "muidea.com/magicCenter/common"

// ID 模块ID
const ID = common.FileRegistryModuleID

// Name 模块名称
const Name = "Magic FileRegistry"

// Description 模块描述信息
const Description = "Magic 文件管理器"

// URL 模块Url
const URL = "/fileregistry"

// PostFile 上传文件
const PostFile = "/file/"

// GetFile 下载文件
const GetFile = "/file/:id"

// DeleteFile 删除文件
const DeleteFile = "/file/:id"
