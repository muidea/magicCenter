package loader

import (
	"log"
	"magiccenter/extern/modules/blog"
	"magiccenter/extern/modules/cms"
	"magiccenter/extern/modules/static"
)

// LoadAllModules 加载所有Module
func LoadAllModules() {
	log.Println("load all modules...")

	blog.LoadModule()

	cms.LoadModule()

	static.LoadModule()
}
