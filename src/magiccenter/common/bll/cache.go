package bll

import "magiccenter/module"

/*
提供临时缓存功能，允许临时保存指定时效的数据
*/

// CacheModuleID Cache 模块ID
const CacheModuleID = "0424492f-420a-42fb-9106-3882c07bf99e"

// InCacheParam 投放数据请求
type InCacheParam struct {
	// Data 放入Cache数据
	Data interface{}
	// MaxAge 数据的最大存档期限，单位为minute
	MaxAge float64
}

// InCacheResult 投放数据结果
type InCacheResult struct {
	// ID Cache分配用于访问数据ID
	ID string
}

// OutCacheParam 获取数据请求
type OutCacheParam struct {
	// ID 访问数据的ID
	ID string
}

// OutCacheResult 获取数据结果
type OutCacheResult struct {
	// Data 获取到的数据
	Data interface{}
	// Found 是否找到数据
	Found bool
}

// RemoveCacheParam 移除缓存数据请求
type RemoveCacheParam struct {
	// ID 访问数据的ID
	ID string
}

// ClearAllCacheParam 清除所有缓存请求
type ClearAllCacheParam struct {
}

// PutInCache 讲数据存入Cache
// maxAge 单位是minute
func PutInCache(data interface{}, maxAge float64) (string, bool) {
	cacheModule, found := module.FindModule(CacheModuleID)
	if !found {
		panic("can't find cache module")
	}

	param := InCacheParam{}
	param.Data = data
	param.MaxAge = maxAge

	result := InCacheResult{}

	if cacheModule.Invoke(&param, &result) {
		return result.ID, true
	}

	return result.ID, false
}

// FetchOutCache 取出数据
func FetchOutCache(id string) (interface{}, bool) {
	cacheModule, found := module.FindModule(CacheModuleID)
	if !found {
		panic("can't find cache module")
	}

	param := OutCacheParam{}
	param.ID = id

	result := OutCacheResult{}

	if cacheModule.Invoke(&param, &result) {
		return result.Data, result.Found
	}

	return result.Data, false
}

// RemoveCache 取出数据
func RemoveCache(id string) {
	cacheModule, found := module.FindModule(CacheModuleID)
	if !found {
		panic("can't find cache module")
	}

	param := RemoveCacheParam{}
	param.ID = id

	cacheModule.Invoke(&param, nil)
}

// ClearAllCache 清空Cache
func ClearAllCache(id string) {
	cacheModule, found := module.FindModule(CacheModuleID)
	if !found {
		panic("can't find cache module")
	}

	param := ClearAllCacheParam{}

	cacheModule.Invoke(&param, nil)
}
