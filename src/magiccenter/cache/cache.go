package cache

import (
	"magiccenter/cache/memorycache"
)

const (
	MEMORY_CACHE int = iota
)

type Cache interface {
	PutIn(data interface{}, maxAge float64) string
	FetchOut(id string) (interface{}, bool)
	ClearAll()
	Release()	
}

var _cache Cache = nil

// 创建指定类型的Cache
func CreateCache( cacheType int ) bool {
	switch cacheType {
		case MEMORY_CACHE:
		_cache = memorycache.NewCache()
		default:
		
	}
	
	return _cache != nil
}

// 销毁Cache
func DestroyCache() {
	if _cache != nil {
		_cache.Release()
		_cache = nil
	}
}

func GetCache() (Cache, bool) {
	if _cache != nil {
		return _cache, true
	}
	
	return nil, false
}

