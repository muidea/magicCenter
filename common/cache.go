package common

// CacheHandler 缓存处理器
type CacheHandler interface {
	PutIn(data interface{}, maxAge float64) string
	FetchOut(id string) (interface{}, bool)
	Remove(id string)
	ClearAll()
}
