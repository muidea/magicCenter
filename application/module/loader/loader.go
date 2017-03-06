package loader

import (
	"muidea.com/magicCenter/application/common/service"
	kernelloader "muidea.com/magiccenter/application/module/kernel/loader"
)

//	externloader "muidea.com/magiccenter/application/module/extern/loader"

// Impl ModuleLoader
type impl struct {
}

// CreateLoader 创建ModuleLader
func CreateLoader() service.ModuleLoader {
	impl := impl{}

	return &impl
}

// LoadAllModules 加载所有Module
func (instance *impl) LoadAllModules(sys service.System) {
	kernelloader.LoadAllModules(sys)
	// externloader.LoadAllModules()
}
