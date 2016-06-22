package cache_test

import (
	"testing"

	"cache"
)

func TestCache(t *testing.T) {
	ch := cache.CreateCache(cache.MEMORY_CACHE)
	if nil == ch {
		t.Error("create cache failed")
	}
}
