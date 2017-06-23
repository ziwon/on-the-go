package cache

import (
	"container/list"
	"sync"
)

type Cache interface {
	Get(keys ...string) []Item

	Put(key string, item Item)

	Remove(keys ...string)

	Size() uint64
}

type Item interface {
	Size() uint64
}

type cached struct {
	item    Item
	element *list.Element
}

func (c *cached) setElementIfNotNil(element *list.Element) {
	if element != nil {
		c.element = element
	}
}

type cache struct {
	sync.Mutex
	cap          uint64
	size         uint64
	items        map[string]*cached
	keyList      *list.List
	recordAdd    func(key string) *list.Element
	recordAccess func(key string) *list.Element
}

type CacheOption func(*cache)

type Policy uint8

const (
	LeastRecentlyAdded Policy = iota
	LeastRecentlyUsed
)

func EvictionPolicy(policy Policy) CacheOption {
	return func(c *cache) {
		switch policy {
		case LeastRecentlyAdded:
			c.recordAccess = c.noop
			c.recordAdd = c.record
		case LeastRecentlyUsed:
			c.recordAccess = c.record
			c.recordAdd = c.noop
		}
	}
}

func New(capacity uint64, options ...CacheOption) Cache {
	c := &cache{
		cap:     capacity,
		keyList: list.New(),
		items:   map[string]*cached{},
	}
	EvictionPolicy(LeastRecentlyUsed)(c)

	for _, option := range options {
		option(c)
	}

	return c
}

func (c *cache) Get(keys ...string) []Item {
	c.Lock()
	defer c.Unlock()

	items := make([]Item, len(keys))
	for i, key := range keys {
		cached := c.items[key]
		if cached == nil {
			items[i] = nil
		} else {
			c.recordAccess(key)
			items[i] = cached.item
		}
	}

	return items
}

func (c *cache) Put(key string, item Item) {
	c.Lock()
	defer c.Unlock()

	c.remove(key)

	c.ensureCapacity(item.Size())

	cached := &cached{item: item}
	cached.setElementIfNotNil(c.recordAdd(key))
	cached.setElementIfNotNil(c.recordAccess(key))
	c.items[key] = cached
	c.size += item.Size()
}

func (c *cache) Remove(keys ...string) {
	c.Lock()
	defer c.Unlock()

	for _, key := range keys {
		c.remove(key)
	}
}

func (c *cache) Size() uint64 {
	return c.size
}

func (c *cache) ensureCapacity(toAdd uint64) {
	mustRemove := int64(c.size+toAdd) - int64(c.cap)
	for mustRemove > 0 {
		key := c.keyList.Back().Value.(string)
		mustRemove -= int64(c.items[key].item.Size())
		c.remove(key)
	}
}

func (c *cache) remove(key string) {
	if cached, ok := c.items[key]; ok {
		delete(c.items, key)
		c.size -= cached.item.Size()
		c.keyList.Remove(cached.element)
	}
}

func (c *cache) noop(string) *list.Element { return nil }

func (c *cache) record(key string) *list.Element {
	if item, ok := c.items[key]; ok {
		c.keyList.MoveToFront(item.element)
		return item.element
	}
	return c.keyList.PushFront(key)
}
