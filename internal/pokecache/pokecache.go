package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct{
    createdAt time.Time
    val []byte
}

type Cache struct{
    cacheEntry map[string]CacheEntry
    mutex sync.Mutex
    interval time.Duration
}

func NewCache(interval time.Duration) *Cache{
    cache := &Cache{
        cacheEntry: make(map[string]CacheEntry),
        mutex: sync.Mutex{},
        interval: interval,
    }
    go cache.reapLoop()
    return cache
}

func (c *Cache) Add(key string, val []byte){
    c.mutex.Lock()

    defer c.mutex.Unlock()

    c.cacheEntry[key] = CacheEntry{
        createdAt: time.Now(),
        val: val,
    } 
}

func (c *Cache) Get(key string) ([]byte, bool){
    c.mutex.Lock()

    defer c.mutex.Unlock()
    entry, found := c.cacheEntry[key]
    if !found {
        return nil, found
    }
    return entry.val,found
}

func (c *Cache) reapLoop(){
    ticker := time.NewTicker(c.interval)
    for range ticker.C{
        c. mutex.Lock()

        now := time.Now()
        for key, entry := range c.cacheEntry{
            if now.Sub(entry.createdAt) > c.interval{
                delete(c.cacheEntry, key)
            }
        }
        c.mutex.Unlock()
    }
}
