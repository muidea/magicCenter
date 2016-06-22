package cache

import (
	"magiccenter/cache/memorycache"
)

const (
	MEMORY_CACHE int = iota
)

// Cache对象，由于系统临时保存信息
// Cache会返回一个string用于应用来获取临时保存的对象
//  存放的对象是有生命周期的，超过设定的存放时间会被系统清除掉
// maxAge 单位为minute
type Cache interface {
	PutIn(data interface{}, maxAge float64) string
	FetchOut(id string) (interface{}, bool)
	Remove(id string)
	ClearAll()
	Release()
}

var _cache Cache = nil
var _cacheType int = 0

// 创建指定类型的Cache
func CreateCache(cacheType int) bool {
	switch cacheType {
	case MEMORY_CACHE:
		_cacheType = MEMORY_CACHE
		_cache = memorycache.NewCache()
	default:

	}

	return _cache != nil
}

// 销毁Cache
func DestroyCache() {
	if _cache != nil {
		_cache.Release()
		switch _cacheType {
		case MEMORY_CACHE:
			memorycache.DestroyCache(_cache)
		}
	}

}

// 获取Cache
func GetCache() (Cache, bool) {
	if _cache != nil {
		return _cache, true
	}

	return nil, false
}
