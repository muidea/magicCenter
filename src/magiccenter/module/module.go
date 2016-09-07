package module

import (
	"log"
	"magiccenter/common"
	"magiccenter/configuration"
	"magiccenter/router"

	"muidea.com/util"
)

// ID -> Module
var moduleIDMap = map[string]common.Module{}

// QueryAllModule 查询所有的模块
// 包含启用和未启用的
func QueryAllModule() []common.Module {
	modules := []common.Module{}

	for _, m := range moduleIDMap {
		modules = append(modules, m)
	}

	return modules
}

// GetAllModuleGroups 获取所有的模块分组
func GetAllModuleGroups() []string {
	allGroups := []string{}
	for _, m := range moduleIDMap {
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
func GetModulesByGroup(group string) []common.Module {
	modules := []common.Module{}
	for _, m := range moduleIDMap {
		g := m.Group()

		if g == group {
			modules = append(modules, m)
		}
	}

	return modules
}

// FindModule 根据Module ID查找指定模块
func FindModule(id string) (common.Module, bool) {
	m, found := moduleIDMap[id]

	return m, found
}

// RegisterModule 在系统中注册模块
func RegisterModule(m common.Module) {
	log.Printf("register module, id:%s, name:%s", m.ID(), m.Name())

	moduleIDMap[m.ID()] = m
}

// UnregisterModule 在系统中取消注册模块
func UnregisterModule(id string) {
	log.Printf("unregister module, id:%s", id)

	delete(moduleIDMap, id)
}

// StartupAllModules 启动全部模块
func StartupAllModules() {
	log.Println("StartupAllModules all modules")

	defaultModule, _ := configuration.GetOption(configuration.SysDefaultModule)

	for _, m := range moduleIDMap {

		routes := m.Routes()
		for _, rt := range routes {
			pattern := "/"
			pattern += util.JoinURL(m.URL(), rt.Pattern())
			if m.ID() == defaultModule {
				pattern = rt.Pattern()
			}

			if rt.Type() == common.GET {
				router.AddGetRoute(pattern, rt.Handler(), rt.Verifier())
			} else if rt.Type() == common.PUT {
				router.AddPutRoute(pattern, rt.Handler(), rt.Verifier())
			} else if rt.Type() == common.POST {
				router.AddPostRoute(pattern, rt.Handler(), rt.Verifier())
			} else if rt.Type() == common.DELETE {
				router.AddDeleteRoute(pattern, rt.Handler(), rt.Verifier())
			} else {
				panic("illegal route type, type:" + rt.Type())
			}
		}

		m.Startup()
	}
}

// CleanupAllModules 清除全部模块
func CleanupAllModules() {
	defaultModule, _ := configuration.GetOption(configuration.SysDefaultModule)

	for _, m := range moduleIDMap {

		routes := m.Routes()
		for _, rt := range routes {
			pattern := util.JoinURL(m.URL(), rt.Pattern())
			if m.ID() == defaultModule {
				pattern = rt.Pattern()
			}

			if rt.Type() == common.GET {
				router.RemoveGetRoute(pattern)
			} else if rt.Type() == common.PUT {
				router.RemovePutRoute(pattern)
			} else if rt.Type() == common.POST {
				router.RemovePostRoute(pattern)
			} else if rt.Type() == common.DELETE {
				router.RemoveDeleteRoute(pattern)
			} else {
				panic("illegal route type, type:" + rt.Type())
			}
		}
		m.Cleanup()
	}
}
