package application

import (
	"os"

	"muidea.com/magicCenter/common/configuration"
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/kernel"
	"muidea.com/magicCenter/module/loader"
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
func AppInstance(server, name, account, password string) Application {
	if app == nil {
		app = &application{}

		app.construct(server, name, account, password)
	}

	return app
}

func (instance *application) construct(server, name, account, password string) {
	os.Setenv("PORT", serverPort)

	dbhelper.InitDB(server, name, account, password)
}

func (instance *application) Run() {
	loader := loader.CreateLoader()
	configuration := configuration.GetSystemConfiguration()

	sys := kernel.NewKernel(loader, configuration)
	sys.StartUp()
	sys.Run()
	sys.ShutDown()
}
