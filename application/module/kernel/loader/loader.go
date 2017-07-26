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
func LoadAllModules(configuration common.Configuration, sessionRegistry common.SessionRegistry, modulHub common.ModuleHub) {

	mail.LoadModule(configuration, modulHub)

	cache.LoadModule(configuration, modulHub)

	account.LoadModule(configuration, modulHub)

	content.LoadModule(configuration, sessionRegistry, modulHub)

	authority.LoadModule(configuration, sessionRegistry, modulHub)

	cas.LoadModule(configuration, sessionRegistry, modulHub)

	static.LoadModule(configuration, modulHub)

	fileregistry.LoadModule(configuration, modulHub)
}
