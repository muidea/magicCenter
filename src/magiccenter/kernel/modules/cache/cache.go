package cache

import (
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/module"

	"muidea.com/util"
)

// ID 模块ID
const ID = "0424492f-420a-42fb-9106-3882c07bf99e"

// Name 块名称
const Name = "Magic Cache"

// Description 模块描述信息
const Description = "Magic 缓存模块"

// URL 模块Url
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
	return common.KERNEL
}

func (instance *cache) URL() string {
	return URL
}

func (instance *cache) EndPoint() common.EndPoint {
	return nil
}

// Route Cache 路由信息
func (instance *cache) Routes() []common.Route {
	routes := []common.Route{}

	return routes
}

// Startup 启动Cache模块
func (instance *cache) Startup() bool {
	return CreateCache(MemoryCache)
}

// Cleanup 清除Cache模块
func (instance *cache) Cleanup() {
	cache, found := GetCache()
	if found {
		cache.Release()

		DestroyCache()
	}
}

// Invoke 执行外部命令
func (instance *cache) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	cache, found := GetCache()
	if !found {
		return false
	}

	switch param.(type) {
	case *commonbll.InCacheParam:
		inBox := param.(*commonbll.InCacheParam)
		if inBox != nil {
			result.(*commonbll.InCacheResult).ID = cache.PutIn(inBox.Data, inBox.MaxAge)
			return true
		}
	case *commonbll.OutCacheParam:
		outBox := param.(*commonbll.OutCacheParam)
		val := result.(*commonbll.OutCacheResult)
		if outBox != nil {
			val.Data, val.Found = cache.FetchOut(outBox.ID)
			return true
		}
	case *commonbll.RemoveCacheParam:
		removeBox := param.(*commonbll.RemoveCacheParam)
		if removeBox != nil {
			cache.Remove(removeBox.ID)
			return true
		}
	case *commonbll.ClearAllCacheParam:
		clearAllBox := param.(*commonbll.ClearAllCacheParam)
		if clearAllBox != nil {
			cache.ClearAll()
			return true
		}
	}
	return false
}
