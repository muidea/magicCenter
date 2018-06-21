package main

import (
	"flag"
	"log"

	"muidea.com/magicCenter/common/initializer"

	"muidea.com/magicCenter"
)

var bindPort = "8080"
var databaseServer = "localhost:3306"
var databaseName = "magiccenter_db"
var databaseAccount = "magiccenter"
var databasePassword = "magiccenter"

func main() {
	flag.StringVar(&bindPort, "ListenPort", bindPort, "magicCenter listen port.")
	flag.StringVar(&databaseServer, "DBServer", databaseServer, "database server address.")
	flag.StringVar(&databaseName, "DBName", databaseName, "database name.")
	flag.StringVar(&databaseAccount, "Account", databaseAccount, "database account.")
	flag.StringVar(&databasePassword, "Password", databasePassword, "database password.")
	flag.Parse()

	log.Println("MagicCenter V1.0")

	app := application.AppInstance(bindPort, databaseServer, databaseName, databaseAccount, databasePassword)

	initializer.InvokHandler()

	app.Run()
}
