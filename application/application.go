package application

import (
	"os"

	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/kernel"
	"muidea.com/magicCenter/application/module/loader"
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

func (instance *application) construct() {
	os.Setenv("PORT", serverPort)
}

func (instance *application) Run() {
	loader := loader.CreateLoader()
	configuration := configuration.CreateConfiguration()

	sys := kernel.NewKernel(loader, configuration)
	sys.StartUp()
	sys.Run()
	sys.ShutDown()
}
