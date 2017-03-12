package api

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/application/module/kernel/api/route"
)

// ID 模块ID
const ID = "5fa671dc-ccb5-4005-8500-f0e45b13705b"

// Name 模块名称
const Name = "Magic Dashboard API"

// Description 模块描述信息
const Description = "Magic Dashboard API模块"

// URL 模块Url
const URL string = "/api"

type api struct {
	moduleHub modulehub.ModuleHub
	routes    []common.Route
}

// LoadModule 加载模块
func LoadModule(cfg configuration.Configuration, modHub modulehub.ModuleHub) {
	instance := &api{moduleHub: modHub, routes: []common.Route{}}

	instance.routes = append(instance.routes, route.CreateGetArticleRoute(modHub))

	modHub.RegisterModule(instance)
}

func (instance *api) ID() string {
	return ID
}

func (instance *api) Name() string {
	return Name
}

func (instance *api) Description() string {
	return Description
}

func (instance *api) Group() string {
	return "admin api"
}

func (instance *api) Type() int {
	return common.KERNEL
}

func (instance *api) URL() string {
	return URL
}

func (instance *api) Status() int {
	return 0
}

func (instance *api) EndPoint() interface{} {
	return nil
}

func (instance *api) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route 路由信息
func (instance *api) Routes() []common.Route {
	return instance.routes
}

// Startup 启动模块
func (instance *api) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *api) Cleanup() {

}
