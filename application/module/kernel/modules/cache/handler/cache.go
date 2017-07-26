package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/foundation/cache"
)

// CreateCacheHandler 新建缓存处理器
func CreateCacheHandler(config common.Configuration, modHub common.ModuleHub) common.CacheHandler {
	return &impl{cache: cache.NewCache()}
}

type impl struct {
	cache cache.Cache
}

func (s *impl) PutIn(data interface{}, maxAge float64) string {
	return s.cache.PutIn(data, maxAge)
}

func (s *impl) FetchOut(id string) (interface{}, bool) {
	return s.cache.FetchOut(id)
}

func (s *impl) Remove(id string) {
	s.cache.Remove(id)
}

func (s *impl) ClearAll() {
	s.cache.ClearAll()
}
