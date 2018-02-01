package model

// AuthGroup 授权组
type AuthGroup struct {
	ID          int
	Name        string
	Description string
}

// ACL 访问控制列表
type ACL struct {
	ID        int
	URL       string
	Method    string
	Module    string
	Status    int
	AuthGroup int
}

// ModuleAuthGroup 模块授权组
type ModuleAuthGroup struct {
	Module    string
	AuthGroup int
}

// UserAuthGroup 用户授权组
type UserAuthGroup struct {
	User      int
	AuthGroup int
}

// UserModuleAuthGroupInfo 用户模块授权组信息
type UserModuleAuthGroupInfo struct {
	User             int
	ModuleAuthGroups []ModuleAuthGroup
}

// ModuleUserAuthGroupInfo 模块用户授权信息
type ModuleUserAuthGroupInfo struct {
	Module         string
	UserAuthGroups []UserAuthGroup
}
