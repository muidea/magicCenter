package loader

import "magiccenter/common"

//externloader "magiccenter/extern/loader"
//kernelloader "magiccenter/kernel/loader"

// Impl ModuleLoader
type impl struct {
}

// CreateLoader 创建ModuleLader
func CreateLoader() common.ModuleLoader {
	impl := impl{}

	return &impl
}

// LoadAllModules 加载所有Module
func (instance impl) LoadAllModules() {
	//kernelloader.LoadAllModules()
	//externloader.LoadAllModules()
}
