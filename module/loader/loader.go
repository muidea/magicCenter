package loader

import (
	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/account"
	"github.com/muidea/magicCenter/module/modules/authority"
	"github.com/muidea/magicCenter/module/modules/cache"
	"github.com/muidea/magicCenter/module/modules/cas"
	"github.com/muidea/magicCenter/module/modules/content"
	"github.com/muidea/magicCenter/module/modules/endpoint"
	"github.com/muidea/magicCenter/module/modules/fileregistry"
	"github.com/muidea/magicCenter/module/modules/mail"
	"github.com/muidea/magicCenter/module/modules/module"
	"github.com/muidea/magicCenter/module/modules/static"
	"github.com/muidea/magicCenter/module/modules/system"
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

	fileregistry.LoadModule(configuration, sessionRegistry, moduleHub)

	static.LoadModule(configuration, sessionRegistry, moduleHub)

	system.LoadModule(configuration, sessionRegistry, moduleHub)

	module.LoadModule(configuration, sessionRegistry, moduleHub)

	mail.LoadModule(configuration, sessionRegistry, moduleHub)

	cache.LoadModule(configuration, sessionRegistry, moduleHub)

	account.LoadModule(configuration, sessionRegistry, moduleHub)

	content.LoadModule(configuration, sessionRegistry, moduleHub)

	endpoint.LoadModule(configuration, sessionRegistry, moduleHub)

	cas.LoadModule(configuration, sessionRegistry, moduleHub)

	authority.LoadModule(configuration, sessionRegistry, moduleHub)

}
