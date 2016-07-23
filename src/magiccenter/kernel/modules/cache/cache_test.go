package cache_test

import (
	"magiccenter/cache"
	"testing"
)

func TestCache(t *testing.T) {
	ch := cache.CreateCache(cache.MemoryCache)
	if !ch {
		t.Error("create cache failed")
	}
}
