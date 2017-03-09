/*
Module管理

提供系统module注册，注销，启动，清除，查找功能

系统中所有的Module在加载完成后都会注册到这里，进行管理

*/

package modulehub

import (
	"log"

	"muidea.com/magicCenter/application/common"
)

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

// impl ModuleHub 实现
type impl struct {
	// ID -> Module
	moduleIDMap map[string]common.Module
}

// CreateModuleHub 创建ModuleHub
func CreateModuleHub() ModuleHub {
	impl := impl{}
	impl.moduleIDMap = map[string]common.Module{}

	return &impl
}

// QueryAllModule 查询所有的模块
// 包含启用和未启用的
func (instance *impl) QueryAllModule() []common.Module {
	modules := []common.Module{}

	for _, m := range instance.moduleIDMap {
		modules = append(modules, m)
	}

	return modules
}

// GetAllModuleGroups 获取所有的模块分组
func (instance *impl) GetAllModuleGroups() []string {
	allGroups := []string{}
	for _, m := range instance.moduleIDMap {
		g := m.Group()

		found := false
		for _, c := range allGroups {
			if g == c {
				found = true
			}
		}
		if !found {
			allGroups = append(allGroups, g)
		}
	}

	return allGroups
}

// GetModulesByGroup 获取指定分组的所有模块
func (instance *impl) GetModulesByGroup(group string) []common.Module {
	modules := []common.Module{}
	for _, m := range instance.moduleIDMap {
		g := m.Group()

		if g == group {
			modules = append(modules, m)
		}
	}

	return modules
}

// FindModule 根据Module ID查找指定模块
func (instance *impl) FindModule(id string) (common.Module, bool) {
	m, found := instance.moduleIDMap[id]

	return m, found
}

// RegisterModule 在系统中注册模块
func (instance *impl) RegisterModule(m common.Module) {
	log.Printf("register module, id:%s, name:%s", m.ID(), m.Name())

	instance.moduleIDMap[m.ID()] = m
}

// UnregisterModule 在系统中取消注册模块
func (instance *impl) UnregisterModule(id string) {
	log.Printf("unregister module, id:%s", id)

	delete(instance.moduleIDMap, id)
}

// StartupAllModules 启动全部模块
func (instance *impl) StartupAllModules() {
	log.Println("StartupAllModules all modules")

	for _, m := range instance.moduleIDMap {

		m.Startup()
	}
}

// CleanupAllModules 清除全部模块
func (instance *impl) CleanupAllModules() {
	for _, m := range instance.moduleIDMap {
		m.Cleanup()
	}
}
