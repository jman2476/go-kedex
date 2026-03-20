package pokecache

import "time"

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = newEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]

	if !ok {
		return nil, ok
	}

	return entry.val, ok
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		for key, entry := range c.entries {
			t := <-ticker.C
			if t.Sub(entry.createdAt) > c.interval {
				c.mu.Lock()
				defer c.mu.Unlock()
				delete(c.entries, key)
			}
		}
	}
}
