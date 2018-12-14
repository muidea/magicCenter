package common

// CacheHandler 缓存处理器
type CacheHandler interface {
	Put(data interface{}, maxAge float64) string
	Fetch(id string) (interface{}, bool)
	Remove(id string)
	ClearAll()
}
