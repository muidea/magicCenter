package service


// ModuleLoader Module加载器
type ModuleLoader interface {
	LoadAllModules(sys System)
}
