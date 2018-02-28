package model

// AuthGroup 授权组
type AuthGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ACL 访问控制列表
type ACL struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Method    string `json:"method"`
	Module    string `json:"module"`
	Status    int    `json:"status"`
	AuthGroup int    `json:"authGroup"`
}

// ModuleAuthGroup 模块授权组
type ModuleAuthGroup struct {
	Module    string `json:"module"`
	AuthGroup int    `json:"authGroup"`
}

// UserAuthGroup 用户授权组
type UserAuthGroup struct {
	User      int `json:"user"`
	AuthGroup int `json:"authGroup"`
}

// UserModuleAuthGroupInfo 用户模块授权组信息
type UserModuleAuthGroupInfo struct {
	User             int               `json:"user"`
	ModuleAuthGroups []ModuleAuthGroup `json:"moduleAuthGroups"`
}

// ModuleUserAuthGroupInfo 模块用户授权信息
type ModuleUserAuthGroupInfo struct {
	Module         string          `json:"module"`
	UserAuthGroups []UserAuthGroup `json:"userAuthGroups"`
}
