package blog

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/extern/modules/blog/def"
	"muidea.com/magicCenter/application/module/extern/modules/blog/route"
)

// LoadModule 加载模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {

	mod, _ := moduleHub.FindModule(common.CotentModuleID)
	contentHandler := mod.EntryPoint().(common.ContentHandler)

	instance := &blog{routes: make([]common.Route, 0)}

	instance.routes = route.AppendBlogRoute(instance.routes, contentHandler)

	moduleHub.RegisterModule(instance)
}

type blog struct {
	routes []common.Route
}

func (s *blog) ID() string {
	return def.ID
}

func (s *blog) Name() string {
	return def.Name
}

func (s *blog) Description() string {
	return def.Description
}

func (s *blog) Group() string {
	return "user"
}

func (s *blog) Type() int {
	return common.EXTERNAL
}

func (s *blog) Status() int {
	return 0
}

func (s *blog) EntryPoint() interface{} {
	return nil
}

// Route Account 路由信息
func (s *blog) Routes() []common.Route {
	return s.routes
}

// Startup 启动Account模块
func (s *blog) Startup() bool {
	return true
}

// Cleanup 清除Account模块
func (s *blog) Cleanup() {
}
