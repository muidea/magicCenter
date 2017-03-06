package loader

import (
	"muidea.com/magicCenter/application/common/service"
	"muidea.com/magicCenter/application/module/kernel/modules/mail"
)

// LoadAllModules 加载所有模块
func LoadAllModules(sys service.System) {

	mail.LoadModule(sys)

	//cache.LoadModule()

	//dashboard.LoadModule()

	//account.LoadModule()

	//content.LoadModule()

	//authority.LoadModule()
}
