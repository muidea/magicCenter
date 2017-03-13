package loader

import (
	"muidea.com/magicCenter/application/common"
	kernelloader "muidea.com/magiccenter/application/module/kernel/loader"
)

//	externloader "muidea.com/magiccenter/application/module/extern/loader"

// Impl ModuleLoader
type impl struct {
}

// CreateLoader 创建ModuleLader
func CreateLoader() common.ModuleLoader {
	impl := impl{}

	return &impl
}

// LoadAllModules 加载所有Module
func (instance *impl) LoadAllModules(configuration common.Configuration, sessionRegistry common.SessionRegistry, modulHub common.ModuleHub) {
	kernelloader.LoadAllModules(configuration, sessionRegistry, modulHub)
	// externloader.LoadAllModules()
}
