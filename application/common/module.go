package common

import "muidea.com/magicCenter/application/common/model"

// 模块类型
const (
	// 内核模块，不能被禁用
	KERNEL = iota
	// 内置模块，属于系统自带可选模块，可以被禁用
	INTERNAL
	// 外部模块，通过外部接口注册进来的模块，可以被禁用
	EXTERNAL
)

const (
	// AuthorityModuleID Authority模块ID
	AuthorityModuleID = "158e11b7-adee-4b0d-afc9-0b47145195bd"
	// CASModuleID CAS模块ID
	CASModuleID = "759a2ee4-147a-4169-ba89-15c0c692bc16"
	// CotentModuleID 内容管理模块ID
	CotentModuleID = "3a7123ec-63f0-5e46-1234-e6ca1af6fe4e"
	// AccountModuleID 账号管理模块ID
	AccountModuleID = "b9e35167-b2a3-43ae-8c57-9b4379475e47"
	// FileRegistryModuleID 文件管理模块ID
	FileRegistryModuleID = "b467c59d-10a5-4875-b617-66662f8824fa"
)

// Module 功能模块接口
type Module interface {
	ID() string
	Name() string
	Description() string
	Group() string
	Type() int
	Status() int

	// AuthGroups 授权组信息
	AuthGroups() []model.AuthGroup

	// Routes 模块支持的路由信息
	Routes() []Route

	// EntryPoint 模块提供访问接口
	EntryPoint() interface{}

	//Startup 启动模块
	Startup() bool

	// Cleanup 清除模块
	Cleanup()
}

// ModuleHub 模块管理器
type ModuleHub interface {
	// 注册Module
	RegisterModule(m Module)
	// 注销Module
	UnregisterModule(id string)
	// 启动所有Module
	StartupAllModules()
	// 清理所有Module
	CleanupAllModules()
	// 查询全部Module
	QueryAllModule() []Module
	// 查询全部Module分组
	GetAllModuleGroups() []string
	// 查询指定分组的Module
	GetModulesByGroup(group string) []Module
	// 查询指定Module
	FindModule(id string) (Module, bool)
}
