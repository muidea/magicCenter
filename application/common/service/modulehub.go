package service

import "muidea.com/magicCenter/application/common"

// ModuleHub 模块管理器
type ModuleHub interface {
	// 注册Module
	RegisterModule(m common.Module)
	// 注销Module
	UnregisterModule(id string)
	// 启动所有Module
	StartupAllModules()
	// 清理所有Module
	CleanupAllModules()
	// 查询全部Module
	QueryAllModule() []common.Module
	// 查询全部Module分组
	GetAllModuleGroups() []string
	// 查询指定分组的Module
	GetModulesByGroup(group string) []common.Module
	// 查询指定Module
	FindModule(id string) (common.Module, bool)
}
