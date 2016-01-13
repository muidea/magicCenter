package main
 
import (
    "log"
    "webcenter/application"
    _"webcenter/session"
    _"webcenter/module"
    _"webcenter/admin"
    _"webcenter/admin/common"
    _"webcenter/admin/auth"
    _"webcenter/admin/content"
    _"webcenter/blog"
    _"webcenter/ui"
    
    "webcenter/module"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0")
	
	app := application.AppInstance()
	
	module.StarupAllModules()
	
	app.Run()
	
	module.CleanupAllModules()
}

