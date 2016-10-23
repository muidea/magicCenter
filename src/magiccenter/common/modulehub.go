package common

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
