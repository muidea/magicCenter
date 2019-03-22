package cas

import (
	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/cas/def"
	"github.com/muidea/magicCenter/module/modules/cas/handler"
	"github.com/muidea/magicCenter/module/modules/cas/route"
	common_const "github.com/muidea/magicCommon/common"
)

type cas struct {
	moduleHub       common.ModuleHub
	sessionRegistry common.SessionRegistry
	routes          []common.Route
	casHandler      common.CASHandler
}

// LoadModule 加载模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &cas{
		moduleHub:       moduleHub,
		sessionRegistry: sessionRegistry,
		casHandler:      handler.CreateCASHandler(moduleHub)}

	instance.routes = route.AppendAccountRoute(instance.routes, instance.casHandler, sessionRegistry)

	moduleHub.RegisterModule(instance)
}

func (s *cas) ID() string {
	return def.ID
}

func (s *cas) Name() string {
	return def.Name
}

func (s *cas) Description() string {
	return def.Description
}

func (s *cas) Group() string {
	return "kernel"
}

func (s *cas) Type() int {
	return common_const.KERNEL
}

func (s *cas) Status() int {
	return common_const.ACTIVE
}

func (s *cas) EntryPoint() interface{} {
	return s.casHandler
}

// Route 路由信息
func (s *cas) Routes() []common.Route {
	return s.routes
}

// Startup 启动模块
func (s *cas) Startup() bool {
	return true
}

// Cleanup 清除模块
func (s *cas) Cleanup() {

}
