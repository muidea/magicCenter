package model

// AuthGroup 授权组
type AuthGroup struct {
	Unit
	Description string `json:"description"`
}

// ACL 访问控制
type ACL struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Method string `json:"method"`
}

// ACLView acl
type ACLView struct {
	ACL
}

// ACLDetail 访问控制列表
type ACLDetail struct {
	ACL
	Module    string `json:"module"`
	Status    int    `json:"status"`
	AuthGroup int    `json:"authGroup"`
}

// ACLDetailView ACL显示信息
type ACLDetailView struct {
	ACLDetail
	Module    Module `json:"module"`
	AuthGroup Unit   `json:"authGroup"`
}

// ModuleUserInfo 模块的用户信息
type ModuleUserInfo struct {
	Module string `json:"module"`
	User   []int  `json:"user"`
}

// ModuleUserInfoView 模块的用户信息
type ModuleUserInfoView struct {
	Module Module `json:"module"`
	User   []User `json:"user"`
}

// UserAuthGroup 用户授权组
type UserAuthGroup struct {
	User      int `json:"user"`
	AuthGroup int `json:"authGroup"`
}

// UserAuthGroupView 用户授权组显示信息
type UserAuthGroupView struct {
	User      User `json:"user"`
	AuthGroup Unit `json:"authGroup"`
}

// UserModuleInfo 用户模块信息
type UserModuleInfo struct {
	User   int      `json:"user"`
	Module []string `json:"module"`
}

// UserModuleInfoView 用户模块信息
type UserModuleInfoView struct {
	User   User     `json:"user"`
	Module []Module `json:"module"`
}

// ModuleAuthGroup 模块授权组
type ModuleAuthGroup struct {
	Module    string `json:"module"`
	AuthGroup int    `json:"authGroup"`
}

// ModuleAuthGroupView 模块授权组显示信息
type ModuleAuthGroupView struct {
	Module    Module `json:"module"`
	AuthGroup Unit   `json:"authGroup"`
}
