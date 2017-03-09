package loader

import kernelloader "muidea.com/magiccenter/application/module/kernel/loader"

//	externloader "muidea.com/magiccenter/application/module/extern/loader"

// ModuleLoader Module加载器
type ModuleLoader interface {
	LoadAllModules()
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
func (instance *impl) LoadAllModules() {
	kernelloader.LoadAllModules()
	// externloader.LoadAllModules()
}
