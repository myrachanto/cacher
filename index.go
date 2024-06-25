package cacher

import (
	"fmt"
	"sync"
)

const (
	READ  = "READ"
	WRITE = "WRITE"
)

var (
	once  sync.Once
	cache Cache
)

type CacheInterface interface {
	Put(username, module, key string, right bool)
	Get(username, module, key string) (bool, error)
	Invalidate(username string)
}

type Cache struct {
	locker sync.Mutex
	Store  map[string]map[string]map[string]bool
}

func NewCache() CacheInterface {
	once.Do(func() {
		cache = Cache{
			Store: make(map[string]map[string]map[string]bool),
		}
	})
	return &cache
}
func (c *Cache) Put(username, module, key string, right bool) {
	c.locker.Lock()
	defer c.locker.Unlock()

	if _, exists := c.Store[username]; !exists {
		c.Store[username] = make(map[string]map[string]bool)
	}
	if _, exists := c.Store[username][module]; !exists {
		c.Store[username][module] = make(map[string]bool)
	}
	c.Store[username][module][key] = right
}
func (c *Cache) Get(username, module, key string) (bool, error) {
	c.locker.Lock()
	defer c.locker.Unlock()

	users, ok := c.Store[username]
	if !ok {
		return false, fmt.Errorf("no such user")
	}
	modulesname, ok := users[module]
	if !ok {
		return false, fmt.Errorf("no such module")
	}
	keys, ok := modulesname[key]
	if !ok {
		return false, fmt.Errorf("no such key available include %s and %s", READ, WRITE)
	}
	return keys, nil
}
func (c *Cache) Invalidate(username string) {
	c.locker.Lock()
	delete(c.Store, username)
	c.locker.Unlock()
}
