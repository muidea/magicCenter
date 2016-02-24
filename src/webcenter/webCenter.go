package main
 
import (
    "log"
    "webcenter/application"
)
 
func main() {
	log.Println("Magic Center V1.0")
	
	app := application.AppInstance()
	
	app.Run()	
}

