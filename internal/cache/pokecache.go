package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type CacheMethods interface {
	Get(string) []byte
	Add(string, []byte)
	reapLoop()
}

type Cache struct {
	Data map[string]cacheEntry
	Mu   *sync.RWMutex
}

func (this *Cache) Get(key string) ([]byte, bool) {
	this.Mu.RLock()
	defer this.Mu.RUnlock()

	x, ok := this.Data[key]
	return x.Val, ok
}

func (this *Cache) Add(key string, val []byte) {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	this.Data[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (this *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {
		this.Mu.Lock()

		for key, entry := range this.Data {
			if time.Since(entry.CreatedAt) > interval {
				delete(this.Data, key)
			}
		}
		this.Mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {

	cache := &Cache{
		Data: make(map[string]cacheEntry),
		Mu:   &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache

}
