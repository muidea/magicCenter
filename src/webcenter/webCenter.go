package main
 
import (
    "log"
    "webcenter/application"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0")
	
	app := application.AppInstance()
	
	app.Run()	
}

