package pokecache

import "time"

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
	}
	go cache.reapLoop()

	return cache
}
