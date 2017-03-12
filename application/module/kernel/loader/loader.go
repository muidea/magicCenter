package loader

import (
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/application/module/kernel/api"
	"muidea.com/magicCenter/application/module/kernel/modules/account"
	"muidea.com/magicCenter/application/module/kernel/modules/cache"
	"muidea.com/magicCenter/application/module/kernel/modules/content"
	"muidea.com/magicCenter/application/module/kernel/modules/mail"
)

// LoadAllModules 加载所有模块
func LoadAllModules(configuration configuration.Configuration, modulHub modulehub.ModuleHub) {

	mail.LoadModule(configuration, modulHub)

	cache.LoadModule(configuration, modulHub)

	api.LoadModule(configuration, modulHub)

	//dashboard.LoadModule()

	account.LoadModule(configuration, modulHub)

	content.LoadModule(configuration, modulHub)

	//authority.LoadModule()
}
