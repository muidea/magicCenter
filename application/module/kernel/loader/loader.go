package loader

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/account"
	"muidea.com/magicCenter/application/module/kernel/modules/authority"
	"muidea.com/magicCenter/application/module/kernel/modules/cache"
	"muidea.com/magicCenter/application/module/kernel/modules/cas"
	"muidea.com/magicCenter/application/module/kernel/modules/content"
	"muidea.com/magicCenter/application/module/kernel/modules/fileregistry"
	"muidea.com/magicCenter/application/module/kernel/modules/mail"
	"muidea.com/magicCenter/application/module/kernel/modules/static"
)

// LoadAllModules 加载所有模块
func LoadAllModules(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {

	mail.LoadModule(configuration, sessionRegistry, moduleHub)

	cache.LoadModule(configuration, sessionRegistry, moduleHub)

	account.LoadModule(configuration, sessionRegistry, moduleHub)

	content.LoadModule(configuration, sessionRegistry, moduleHub)

	cas.LoadModule(configuration, sessionRegistry, moduleHub)

	authority.LoadModule(configuration, sessionRegistry, moduleHub)

	static.LoadModule(configuration, sessionRegistry, moduleHub)

	fileregistry.LoadModule(configuration, sessionRegistry, moduleHub)
}
