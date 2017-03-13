package common

// ModuleLoader Module加载器
type ModuleLoader interface {
	LoadAllModules(configuration Configuration, sessionRegstry SessionRegistry, modulHub ModuleHub)
}
