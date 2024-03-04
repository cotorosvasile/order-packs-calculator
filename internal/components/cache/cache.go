package cache

import "sync"

type Cache interface {
	Get() []int
	Set([]int)
}

type cache struct {
	mu    sync.RWMutex
	cache []int
}

func NewCache() Cache {
	return &cache{cache: []int{250, 500, 1000, 2000, 5000}}
}

func (c *cache) Get() []int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.cache
}

func (c *cache) Set(boxSizes []int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache = boxSizes
}
