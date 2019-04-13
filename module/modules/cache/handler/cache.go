package handler

import (
	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCommon/foundation/cache"
)

// CreateCacheHandler 新建缓存处理器
func CreateCacheHandler(config common.Configuration, modHub common.ModuleHub) common.CacheHandler {
	return &impl{cache: cache.NewCache()}
}

type impl struct {
	cache cache.Cache
}

func (s *impl) Put(data interface{}, maxAge float64) string {
	return s.cache.Put(data, maxAge)
}

func (s *impl) Fetch(id string) (interface{}, bool) {
	return s.cache.Fetch(id)
}

func (s *impl) Remove(id string) {
	s.cache.Remove(id)
}

func (s *impl) ClearAll() {
	s.cache.ClearAll()
}
