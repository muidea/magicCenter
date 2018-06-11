package cache

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/cache/def"
	"muidea.com/magicCenter/module/modules/cache/handler"
	"muidea.com/magicCenter/module/modules/cache/route"
	common_const "muidea.com/magicCommon/common"
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
	return common_const.KERNEL
}

func (s *cacheModule) Status() int {
	return common_const.ACTIVE
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
