package main

import "sync"

type Cache struct {
	data map[string]interface{}
	mx   sync.RWMutex
}

func (ch *Cache) Set(key string, value interface{}) {
	defer ch.mx.Unlock()
	ch.mx.Lock()
	ch.data[key] = value
}

func (ch *Cache) Get(key string) interface{} {
	defer ch.mx.RUnlock()
	ch.mx.RLock()
	if _, ok := ch.data[key]; ok {
		return ch.data[key]
	} else {
		return nil
	}
}

func (ch *Cache) Delete(key string) {
	defer ch.mx.Unlock()
	ch.mx.Lock()

	delete(ch.data, key)
}

func NewCache() Cache {
	return Cache{
		data: make(map[string]interface{}),
		mx:   sync.RWMutex{},
	}
}
