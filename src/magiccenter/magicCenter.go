package main
 
import (
    "log"
    "magiccenter/application"
)
 
func main() {
	log.Println("MagicCenter V1.0")
	
	app := application.AppInstance()
	
	app.Run()	
}

