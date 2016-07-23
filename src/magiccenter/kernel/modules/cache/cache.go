package cache

import (
	"magiccenter/module"
	"magiccenter/router"
)

// ID Cache模块ID
const ID = "0168384d-900c-47c0-b5b7-693169141979"

// Name Cache块名称
const Name = "Magic Cache"

// Description Cache模块描述信息
const Description = "Magic 缓存模块"

// URL Cache模块Url
const URL string = "cache"

type cache struct {
}

var instance *cache

// LoadModule 加载Cache模块
func LoadModule() {
	if instance == nil {
		instance = &cache{}
	}

	module.RegisterModule(instance)
}

func (instance *cache) ID() string {
	return ID
}

func (instance *cache) Name() string {
	return Name
}

func (instance *cache) Description() string {
	return Description
}

func (instance *cache) Group() string {
	return "util"
}

func (instance *cache) Type() int {
	return module.KERNEL
}

func (instance *cache) URL() string {
	return URL
}

func (instance *cache) Resource() module.Resource {
	return nil
}

// Route Cache 路由信息
func (instance *cache) Routes() []router.Route {
	routes := []router.Route{}

	return routes
}

// Startup 启动Cache模块
func (instance *cache) Startup() bool {
	return true
}

// Cleanup 清除Cache模块
func (instance *cache) Cleanup() {

}

// Invoke 执行外部命令
func (instance *cache) Invoke(param interface{}) bool {
	return false
}
