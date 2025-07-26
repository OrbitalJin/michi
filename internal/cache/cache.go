package cache

import "sync"


type Cache[K any, V any] struct {
	data *sync.Map
}

func New[K any, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: &sync.Map{},
	}
}


func (c *Cache[K, V]) Store(key K, value V) {
	c.data.Store(key, value)
}


func (c *Cache[K, V]) Load(key K) (V, bool) {
	v, ok := c.data.Load(key)

	if !ok || v == nil {
		var nilValue V
		return nilValue, false
	}

	return v.(V), ok
}


func (c *Cache[K, V]) Delete(key K) {
	c.data.Delete(key)
} 


func (c *Cache[K, V]) Invalidate() {
	c.data.Clear()
}


