package main
 
import (
    "log"
    "webcenter/application"
    _"webcenter/session"
    _"webcenter/common"
    _"webcenter/ui"
    _"webcenter/auth"
    _"webcenter/admin"
    _"webcenter/content"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0")
	
	app := application.AppInstance()
	
	app.Run()
}

