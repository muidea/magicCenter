package loader

import (
	api "magiccenter/kernel/api"
	"magiccenter/kernel/dashboard"
	"magiccenter/kernel/modules/account"
	"magiccenter/kernel/modules/cache"
	"magiccenter/kernel/modules/content"
	"magiccenter/kernel/modules/mail"
)

// LoadAllModules 加载所有模块
func LoadAllModules() {

	mail.LoadModule()

	cache.LoadModule()

	dashboard.LoadModule()

	account.LoadModule()

	content.LoadModule()

	api.LoadModule()
}
