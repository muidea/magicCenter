package cache

import (
	"magiccenter/common"
	"magiccenter/module"

	"muidea.com/util"
)

// ID 模块ID
const ID = "0168384d-900c-47c0-b5b7-693169141979"

// Name 块名称
const Name = "Magic Cache"

// Description 模块描述信息
const Description = "Magic 缓存模块"

// URL 模块Url
const URL string = "cache"

type cache struct {
}

var instance *cache

// InCacheBox 投放数据请求
type InCacheBox struct {
	// Data 放入Cache数据
	Data interface{}
	// MaxAge 数据的最大存档期限，单位为minute
	MaxAge float64
	// ID Cache分配用于访问数据ID
	ID *string
}

// OutCacheBox 获取数据请求
type OutCacheBox struct {
	// ID 访问数据的ID
	ID string
	// Data 获取到的数据
	Data interface{}
	// Found 是否找到数据
	Found *bool
}

// RemoveCacheBox 移除缓存数据请求
type RemoveCacheBox struct {
	// ID 访问数据的ID
	ID string
}

// ClearAllCacheBox 清除所有缓存请求
type ClearAllCacheBox struct {
}

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

	cache, found := GetCache()
	if !found {
		return false
	}

	inBox := param.(*InCacheBox)
	if inBox != nil {
		*(inBox.ID) = cache.PutIn(inBox.Data, inBox.MaxAge)
		return true
	}

	outBox := param.(*OutCacheBox)
	if outBox != nil {
		outBox.Data, *(outBox.Found) = cache.FetchOut(outBox.ID)
		return true
	}

	removeBox := param.(*RemoveCacheBox)
	if removeBox != nil {
		cache.Remove(removeBox.ID)
		return true
	}

	clearAllBox := param.(*ClearAllCacheBox)
	if clearAllBox != nil {
		cache.ClearAll()
		return true
	}

	return false
}
