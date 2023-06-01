package cache

import (
	"sync"
	"time"
)

type cacheItem struct {
	value    interface{}
	deadLine time.Time
}

type Cache struct {
	data map[string]cacheItem
	mx   sync.RWMutex
}

func (ch *Cache) Set(key string, value interface{}, ttl time.Duration) {
	defer ch.mx.Unlock()
	ch.mx.Lock()
	ci := cacheItem{
		value:    value,
		deadLine: time.Now().Add(ttl),
	}
	ch.data[key] = ci
}

func (ch *Cache) Get(key string) interface{} {
	defer ch.mx.RUnlock()
	ch.mx.RLock()
	if _, ok := ch.data[key]; ok {
		if time.Now().Before(ch.data[key].deadLine) {
			return ch.data[key].value
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (ch *Cache) Delete(key string) {
	defer ch.mx.Unlock()
	ch.mx.Lock()

	delete(ch.data, key)
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]cacheItem),
		mx:   sync.RWMutex{},
	}
}
