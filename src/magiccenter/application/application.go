package application

import (
	"magiccenter/auth"
	"magiccenter/configuration"
	"magiccenter/loader"
	"magiccenter/system"
	"os"
)

var serverPort = "8888"

// Application 接口
type Application interface {
	Run()
}

var app *application

type application struct {
}

// AppInstance 返回Application对象
func AppInstance() Application {
	if app == nil {
		app = &application{}

		app.construct()
	}

	return app
}

func (instance application) construct() {
	os.Setenv("PORT", serverPort)
}

func (instance application) Run() {
	loader := loader.CreateLoader()

	authority := auth.CreateAuthority()

	configuration := configuration.CreateConfiguration()

	system.Initialize(loader, authority, configuration)

	system.Run()

	system.Uninitialize()
}
