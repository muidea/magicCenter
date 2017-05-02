package loader

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/account"
	"muidea.com/magicCenter/application/module/kernel/modules/api"
	"muidea.com/magicCenter/application/module/kernel/modules/cache"
	"muidea.com/magicCenter/application/module/kernel/modules/cas"
	"muidea.com/magicCenter/application/module/kernel/modules/content"
	"muidea.com/magicCenter/application/module/kernel/modules/dashboard"
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

	cas.LoadModule(configuration, sessionRegistry, modulHub)

	dashboard.LoadModule(configuration, modulHub)

	// API 必须放在最后，否则找不到对应的Module
	api.LoadModule(configuration, sessionRegistry, modulHub)

	static.LoadModule(configuration, modulHub)

	fileregistry.LoadModule(configuration, modulHub)
}
