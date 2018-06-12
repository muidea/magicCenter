package main

import (
	"flag"
	"log"

	"muidea.com/magicCenter/common/initializer"

	"muidea.com/magicCenter"
)

var databaseServer = "localhost:3306"
var databaseName = "magiccenter_db"
var databaseAccount = "magiccenter"
var databasePassword = "magiccenter"

func main() {
	flag.StringVar(&databaseServer, "server", databaseServer, "database server address.")
	flag.StringVar(&databaseName, "database", databaseName, "database name.")
	flag.StringVar(&databaseAccount, "account", databaseAccount, "database account.")
	flag.StringVar(&databasePassword, "password", databasePassword, "database password.")
	flag.Parse()

	log.Println("MagicCenter V1.0")

	app := application.AppInstance(databaseServer, databaseName, databaseAccount, databasePassword)

	initializer.InvokHandler()

	app.Run()
}
