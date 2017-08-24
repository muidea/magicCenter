package cors

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/cors/def"
	"muidea.com/magicCenter/application/module/kernel/modules/cors/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/cors/route"
)

type cors struct {
	moduleHub       common.ModuleHub
	sessionRegistry common.SessionRegistry
	routes          []common.Route
	corsHandler     common.CorsHandler
}

// LoadModule 加载模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &cors{
		moduleHub:       moduleHub,
		sessionRegistry: sessionRegistry,
		corsHandler:     handler.CreateCorsHandler(configuration)}

	instance.routes = route.AppendCorsRoute(instance.routes, instance.corsHandler)

	moduleHub.RegisterModule(instance)
}

func (s *cors) ID() string {
	return def.ID
}

func (s *cors) Name() string {
	return def.Name
}

func (s *cors) Description() string {
	return def.Description
}

func (s *cors) Group() string {
	return "kernel"
}

func (s *cors) Type() int {
	return common.KERNEL
}

func (s *cors) Status() int {
	return 0
}

func (s *cors) EntryPoint() interface{} {
	return s.corsHandler
}

func (s *cors) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route 路由信息
func (s *cors) Routes() []common.Route {
	return s.routes
}

// Startup 启动模块
func (s *cors) Startup() bool {
	return true
}

// Cleanup 清除模块
func (s *cors) Cleanup() {

}
