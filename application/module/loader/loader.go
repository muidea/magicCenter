package loader

import (
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/kernel/modulehub"
	kernelloader "muidea.com/magiccenter/application/module/kernel/loader"
)

//	externloader "muidea.com/magiccenter/application/module/extern/loader"

// ModuleLoader Module加载器
type ModuleLoader interface {
	LoadAllModules(configuration configuration.Configuration, modulHub modulehub.ModuleHub)
}

// Impl ModuleLoader
type impl struct {
}

// CreateLoader 创建ModuleLader
func CreateLoader() ModuleLoader {
	impl := impl{}

	return &impl
}

// LoadAllModules 加载所有Module
func (instance *impl) LoadAllModules(configuration configuration.Configuration, modulHub modulehub.ModuleHub) {
	kernelloader.LoadAllModules(configuration, modulHub)
	// externloader.LoadAllModules()
}
