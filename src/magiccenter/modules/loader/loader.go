package loader

import (
	"log"
	"magiccenter/modules/blog"
	"magiccenter/modules/cms"
	"magiccenter/modules/static"
)

// LoadAllModules 加载所有Module
func LoadAllModules() {
	log.Println("load all modules...")

	blog.LoadModule()

	cms.LoadModule()

	static.LoadModule()
}
