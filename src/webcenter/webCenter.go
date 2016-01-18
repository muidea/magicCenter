package main
 
import (
    "log"
    "webcenter/application"
    _"webcenter/module"
    _"webcenter/session"
    _"webcenter/admin"
    _"webcenter/admin/common"
    _"webcenter/admin/auth"
    _"webcenter/admin/content"
    _"webcenter/blog"
    _"webcenter/ui"
    
)
 
func main() {
	log.Println("MagicID WebCenter V1.0")
	
	app := application.AppInstance()
	
	app.Run()	
}

