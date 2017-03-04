package cache

import "testing"

func TestCache(t *testing.T) {
	ch := CreateCache(MemoryCache)
	if !ch {
		t.Error("create cache failed")
	}

	DestroyCache()
}
