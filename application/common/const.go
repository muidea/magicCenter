package common

import "muidea.com/magicCenter/application/common/model"

// SessionID 会话ID
const SessionID = "sessionID"

// AuthTokenID 鉴权Token
const AuthTokenID = "authToken"

// VisitorAuthGroup 访客权限组
var VisitorAuthGroup = model.AuthGroup{ID: 0, Name: "访客权限组", Description: "允许查看公开权限的内容，无须登录"}

// UserAuthGroup 用户权限组
var UserAuthGroup = model.AuthGroup{ID: 1, Name: "用户权限组", Description: "允许查看用户权限的内容以及公开权限的内容，要求预先进行登录"}

// MaintainerAuthGroup 维护权限组
var MaintainerAuthGroup = model.AuthGroup{ID: 2, Name: "维护权限组", Description: "允许查看和编辑内容，要求预先进行登录"}
