package cache

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/cache/def"
	"muidea.com/magicCenter/application/module/kernel/modules/cache/handler"
	"muidea.com/magicCenter/application/module/kernel/modules/cache/route"
)

// LoadModule 加载Cache模块
func LoadModule(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) {
	instance := &cacheModule{cacheHandler: handler.CreateCacheHandler(configuration, moduleHub)}

	instance.routes = route.AppendCacheRoute(instance.routes, instance.cacheHandler, sessionRegistry)
	moduleHub.RegisterModule(instance)
}

type cacheModule struct {
	routes       []common.Route
	cacheHandler common.CacheHandler
}

func (s *cacheModule) ID() string {
	return def.ID
}

func (s *cacheModule) Name() string {
	return def.Name
}

func (s *cacheModule) Description() string {
	return def.Description
}

func (s *cacheModule) Group() string {
	return "util"
}

func (s *cacheModule) Type() int {
	return common.KERNEL
}

func (s *cacheModule) Status() int {
	return 0
}

func (s *cacheModule) EntryPoint() interface{} {
	return s.cacheHandler
}

// Route Cache 路由信息
func (s *cacheModule) Routes() []common.Route {

	return s.routes
}

// Startup 启动Cache模块
func (s *cacheModule) Startup() bool {
	return true
}

// Cleanup 清除Cache模块
func (s *cacheModule) Cleanup() {
}
