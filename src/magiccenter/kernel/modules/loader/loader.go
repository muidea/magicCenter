package loader

import (
	"magiccenter/kernel/modules/account"
	"magiccenter/kernel/modules/cache"
	"magiccenter/kernel/modules/content"
	"magiccenter/kernel/modules/mail"
)

// LoadAllModules 加载所有模块
func LoadAllModules() {

	mail.LoadModule()

	cache.LoadModule()

	account.LoadModule()

	content.LoadModule()
}
