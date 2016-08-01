package module

import (
	"log"
	"magiccenter/configuration"
	"magiccenter/router"

	"muidea.com/util"
)

// 模块类型
const (
	// 内核模块，不能被禁用
	KERNEL = iota
	// 内置模块，属于系统自带可选模块，可以被禁用
	INTERNAL
	// 外部模块，通过外部接口注册进来的模块，可以被禁用
	EXTERNAL
)

// Module 功能模块接口
type Module interface {
	ID() string
	Name() string
	Description() string
	Group() string
	Type() int
	// URL 模块Url，每个模块都对应唯一的Url,不带'/'
	URL() string

	// Resource 模块提供的Rest api支持
	Resource() Resource

	// Routes 模块支持的路由信息
	Routes() []router.Route

	//Startup 启动模块
	Startup() bool

	// Cleanup 清除模块
	Cleanup()

	// Invoke 执行指定操作，实际由各个模块具体定义实现
	Invoke(param interface{}) bool
}

// ID -> Module
var moduleIDMap = map[string]Module{}

// QueryAllModule 查询所有的模块
// 包含启用和未启用的
func QueryAllModule() []Module {
	modules := []Module{}

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
func GetModulesByGroup(group string) []Module {
	modules := []Module{}
	for _, m := range moduleIDMap {
		g := m.Group()

		if g == group {
			modules = append(modules, m)
		}
	}

	return modules
}

// FindModule 根据Module ID查找指定模块
func FindModule(id string) (Module, bool) {
	m, found := moduleIDMap[id]

	return m, found
}

// RegisterModule 在系统中注册模块
func RegisterModule(m Module) {
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
			pattern := util.JoinURL(m.URL(), rt.Pattern())
			if m.ID() == defaultModule {
				pattern = rt.Pattern()
			}

			if rt.Type() == router.GET {
				router.AddGetRoute(pattern, rt.Handler(), rt.Verifier())
			} else if rt.Type() == router.PUT {
				router.AddPutRoute(pattern, rt.Handler(), rt.Verifier())
			} else if rt.Type() == router.POST {
				router.AddPostRoute(pattern, rt.Handler(), rt.Verifier())
			} else if rt.Type() == router.DELETE {
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

			if rt.Type() == router.GET {
				router.RemoveGetRoute(pattern)
			} else if rt.Type() == router.PUT {
				router.RemovePutRoute(pattern)
			} else if rt.Type() == router.POST {
				router.RemovePostRoute(pattern)
			} else if rt.Type() == router.DELETE {
				router.RemoveDeleteRoute(pattern)
			} else {
				panic("illegal route type, type:" + rt.Type())
			}
		}
		m.Cleanup()
	}
}
