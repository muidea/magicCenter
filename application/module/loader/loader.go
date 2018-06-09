package loader

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/modules/account"
	"muidea.com/magicCenter/application/module/modules/authority"
	"muidea.com/magicCenter/application/module/modules/cache"
	"muidea.com/magicCenter/application/module/modules/cas"
	"muidea.com/magicCenter/application/module/modules/content"
	"muidea.com/magicCenter/application/module/modules/fileregistry"
	"muidea.com/magicCenter/application/module/modules/mail"
	"muidea.com/magicCenter/application/module/modules/module"
	"muidea.com/magicCenter/application/module/modules/static"
	"muidea.com/magicCenter/application/module/modules/system"
)

// Impl ModuleLoader
type impl struct {
}

// CreateLoader 创建ModuleLader
func CreateLoader() common.ModuleLoader {
	impl := impl{}

	return &impl
}

// LoadAllModules 加载所有Module
func (instance *impl) LoadAllModules(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	system.LoadModule(configuration, sessionRegistry, moduleHub)

	module.LoadModule(configuration, sessionRegistry, moduleHub)

	mail.LoadModule(configuration, sessionRegistry, moduleHub)

	cache.LoadModule(configuration, sessionRegistry, moduleHub)

	account.LoadModule(configuration, sessionRegistry, moduleHub)

	content.LoadModule(configuration, sessionRegistry, moduleHub)

	cas.LoadModule(configuration, sessionRegistry, moduleHub)

	authority.LoadModule(configuration, sessionRegistry, moduleHub)

	fileregistry.LoadModule(configuration, sessionRegistry, moduleHub)

	static.LoadModule(configuration, sessionRegistry, moduleHub)
}
