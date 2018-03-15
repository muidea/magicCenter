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
	AuthGroup Unit `json:"authGroup"`
}

// ModuleAuthGroup 模块授权组
type ModuleAuthGroup struct {
	Module    string `json:"module"`
	AuthGroup int    `json:"authGroup"`
}

// ModuleAuthGroupView 模块授权组显示信息
type ModuleAuthGroupView struct {
	ModuleAuthGroup
	AuthGroup Unit `json:"authGroup"`
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

// UserModuleAuthGroup 用户模块授权组
type UserModuleAuthGroup struct {
	User            int               `json:"user"`
	ModuleAuthGroup []ModuleAuthGroup `json:"moduleAuthGroup"`
}

// UserModuleAuthGroupView 用户模块授权组显示信息
type UserModuleAuthGroupView struct {
	User            User                  `json:"user"`
	ModuleAuthGroup []ModuleAuthGroupView `json:"moduleAuthGroup"`
}

// ModuleUserAuthGroup 模块用户授权信息
type ModuleUserAuthGroup struct {
	Module        string          `json:"module"`
	UserAuthGroup []UserAuthGroup `json:"userAuthGroup"`
}

// ModuleUserAuthGroupView 模块用户授权信息显示信息
type ModuleUserAuthGroupView struct {
	Module        string              `json:"module"`
	UserAuthGroup []UserAuthGroupView `json:"userAuthGroup"`
}
