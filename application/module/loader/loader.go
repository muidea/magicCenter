package loader

import (
	"muidea.com/magiccenter/application/common"
)

//	externloader "muidea.com/magiccenter/application/module/extern/loader"
//	kernelloader "muidea.com/magiccenter/application/module/kernel/loader"

// Impl ModuleLoader
type impl struct {
}

// CreateLoader 创建ModuleLader
func CreateLoader() common.ModuleLoader {
	impl := impl{}

	return &impl
}

// LoadAllModules 加载所有Module
func (instance *impl) LoadAllModules() {
	//kernelloader.LoadAllModules()
	// externloader.LoadAllModules()
}
