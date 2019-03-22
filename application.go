package application

import (
	"github.com/muidea/magicCenter/common/configuration"
	"github.com/muidea/magicCenter/common/daemon"
	"github.com/muidea/magicCenter/kernel"
	"github.com/muidea/magicCenter/module/loader"
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
	defer sys.ShutDown()

	daemon.Start()
	defer daemon.Stop()

	sys.Run()
}
