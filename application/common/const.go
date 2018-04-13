package common

import "muidea.com/magicCommon/model"

// SessionID 会话ID
const SessionID = "sessionID"

// AuthTokenID 鉴权Token
const AuthTokenID = "authToken"

// VisitorAuthGroup 访客权限组
var VisitorAuthGroup = model.AuthGroup{Unit: model.Unit{ID: 0, Name: "访客权限组"}, Description: "允许查看公开权限的内容，无须登录"}

// UserAuthGroup 用户权限组
var UserAuthGroup = model.AuthGroup{Unit: model.Unit{ID: 1, Name: "用户权限组"}, Description: "允许查看用户权限的内容以及公开权限的内容，要求预先进行登录"}

// MaintainerAuthGroup 维护权限组
var MaintainerAuthGroup = model.AuthGroup{Unit: model.Unit{ID: 2, Name: "维护权限组"}, Description: "允许查看和编辑内容，要求预先进行登录"}

// DefaultContentCatalog 系统默认的Content分组，UpdataCatalog时，如果需要创建Catalog,则默认指定的ParentCatalog
var DefaultContentCatalog = model.CatalogDetail{Summary: model.Summary{Unit: model.Unit{ID: 0, Name: "默认Content分组"}, CreateDate: "", Creater: 0}, Description: "系统默认的Content分组"}
