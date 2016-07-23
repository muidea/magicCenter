package cache

import "magiccenter/kernel/modules/cache/memorycache"

const (
	//MemoryCache 内存Cache
	MemoryCache int = iota
)

// Cache 对象，由于系统临时保存信息
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

var _cache Cache
var _cacheType = 0

// CreateCache 创建指定类型的Cache
func CreateCache(cacheType int) bool {
	switch cacheType {
	case MemoryCache:
		_cacheType = MemoryCache
		_cache = memorycache.NewCache()
	default:

	}

	return _cache != nil
}

// DestroyCache 销毁Cache
func DestroyCache() {
	if _cache != nil {
		_cache.Release()
		switch _cacheType {
		case MemoryCache:
			memorycache.DestroyCache(_cache)
		}
	}

}

// GetCache 获取Cache
func GetCache() (Cache, bool) {
	if _cache != nil {
		return _cache, true
	}

	return nil, false
}
