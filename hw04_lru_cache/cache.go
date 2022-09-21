package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	newItem := cacheItem{key: key, value: value}
	c.Lock()
	defer c.Unlock()
	if val, ok := c.items[key]; ok {
		val.Value = newItem
		c.queue.MoveToFront(val)
		return true
	}
	if c.queue.Len() == c.capacity {
		delete(c.items, c.queue.Back().Value.(cacheItem).key)
		c.queue.Remove(c.queue.Back())
	}
	c.items[key] = c.queue.PushFront(newItem)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()
	if val, ok := c.items[key]; ok {
		c.queue.MoveToFront(val)
		return val.Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
