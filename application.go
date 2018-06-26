package application

import (
	"muidea.com/magicCenter/common/configuration"
	"muidea.com/magicCenter/kernel"
	"muidea.com/magicCenter/module/loader"
)

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
	}

	return app
}

func (instance *application) Run() {
	loader := loader.CreateLoader()
	configuration := configuration.GetSystemConfiguration()

	sys := kernel.NewKernel(loader, configuration)
	sys.StartUp()
	sys.Run()
	sys.ShutDown()
}
