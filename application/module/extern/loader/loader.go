package loader

import (
	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/extern/modules/blog"
)

// LoadAllModules 加载所有Module
func LoadAllModules(configuration common.Configuration, sessionRegistry common.SessionRegistry, modulHub common.ModuleHub) {
	log.Println("load all modules...")

	blog.LoadModule(configuration, sessionRegistry, modulHub)

	//cms.LoadModule()

	//static.LoadModule()
}
