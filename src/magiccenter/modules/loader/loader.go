package loader

import (
	"log"
	"magiccenter/modules/blog"
	"magiccenter/modules/blog2"
)

func LoadAllModules() {
	log.Println("load all modules...")
	
	blog.LoadModule()
	
	blog2.LoadModule()
}