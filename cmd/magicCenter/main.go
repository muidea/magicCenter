package main

import (
	"log"

	"muidea.com/magicCenter/application"
)

func main() {
	log.Println("MagicCenter V1.0")

	app := application.AppInstance()

	app.Run()
}
