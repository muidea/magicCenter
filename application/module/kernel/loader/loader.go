package loader

import (
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/application/module/kernel/modules/mail"
)

// LoadAllModules 加载所有模块
func LoadAllModules(configuration configuration.Configuration, modulHub modulehub.ModuleHub) {

	mail.LoadModule(configuration, modulHub)

	//cache.LoadModule()

	//dashboard.LoadModule()

	//account.LoadModule()

	//content.LoadModule()

	//authority.LoadModule()
}
