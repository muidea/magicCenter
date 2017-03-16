package loader

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/api"
	"muidea.com/magicCenter/application/module/kernel/authority"
	"muidea.com/magicCenter/application/module/kernel/modules/account"
	"muidea.com/magicCenter/application/module/kernel/modules/cache"
	"muidea.com/magicCenter/application/module/kernel/modules/content"
	"muidea.com/magicCenter/application/module/kernel/modules/mail"
)

// LoadAllModules 加载所有模块
func LoadAllModules(configuration common.Configuration, sessionRegistry common.SessionRegistry, modulHub common.ModuleHub) {

	mail.LoadModule(configuration, modulHub)

	cache.LoadModule(configuration, modulHub)

	//dashboard.LoadModule()

	account.LoadModule(configuration, modulHub)

	content.LoadModule(configuration, modulHub)

	authority.LoadModule(configuration, sessionRegistry, modulHub)

	// API 必须放在最后，否则找不到对应的Module
	api.LoadModule(configuration, sessionRegistry, modulHub)
}
