package api

import (
	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
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
	moduleHub       common.ModuleHub
	sessionRegistry common.SessionRegistry
	routes          []common.Route
}

// LoadModule 加载模块
func LoadModule(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) {
	instance := &api{moduleHub: modHub, sessionRegistry: sessionRegistry, routes: []common.Route{}}

	rt, ok := route.CreateGetArticleRoute(modHub)
	if ok {
		instance.routes = append(instance.routes, rt)
	} else {
		log.Print("CreateGetArticleRoute failed")
	}

	rt, ok = route.CreateGetAllArticleRoute(modHub)
	if ok {
		instance.routes = append(instance.routes, rt)
	} else {
		log.Print("CreateGetAllArticleRoute failed")
	}

	rt, ok = route.CreateGetByCatalogArticleRoute(modHub)
	if ok {
		instance.routes = append(instance.routes, rt)
	} else {
		log.Print("CreateGetByCatalogArticleRoute failed")
	}

	rt, ok = route.CreateCreateArticleRoute(modHub, sessionRegistry)
	if ok {
		instance.routes = append(instance.routes, rt)
	} else {
		log.Print("CreateCreateArticleRoute failed")
	}
	rt, ok = route.CreateUpdateArticleRoute(modHub, sessionRegistry)
	if ok {
		instance.routes = append(instance.routes, rt)
	} else {
		log.Print("CreateUpdateArticleRoute failed")
	}
	rt, ok = route.CreateDestroyArticleRoute(modHub, sessionRegistry)
	if ok {
		instance.routes = append(instance.routes, rt)
	} else {
		log.Print("CreateDestroyArticleRoute failed")
	}
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
