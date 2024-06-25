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
	Get(username, module, key string, right bool) (bool, error)
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
	rights := map[string]bool{key: right}
	modulesname := map[string]map[string]bool{module: rights}
	c.Store[username] = modulesname
	c.locker.Unlock()
}
func (c *Cache) Get(username, module, key string, right bool) (bool, error) {
	c.locker.Lock()
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
		return false, fmt.Errorf("no such key avaliable include %s and %s", READ, WRITE)
	}
	c.locker.Unlock()
	return keys, nil
}
func (c *Cache) Invalidate(username string) {
	c.locker.Lock()
	delete(c.Store, username)
	c.locker.Unlock()
}
