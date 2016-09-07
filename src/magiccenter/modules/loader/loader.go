package loader

import (
	"log"
	"magiccenter/kernel/modules/loader"
	"magiccenter/modules/blog"
	"magiccenter/modules/cms"
	"magiccenter/modules/static"
)

// LoadAllModules 加载所有Module
func LoadAllModules() {
	log.Println("load all modules...")

	loader.LoadAllModules()

	blog.LoadModule()

	cms.LoadModule()

	static.LoadModule()
}
