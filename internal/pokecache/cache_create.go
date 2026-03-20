package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
		mu:       &sync.Mutex{},
		entries:  map[string]cacheEntry{},
	}
	go cache.reapLoop()

	return cache
}
