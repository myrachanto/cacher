package cacher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheWorking(t *testing.T) {
	cache := NewCache()
	user := "myrachanto"
	module := "users"
	key := "READ"
	right := true
	cache.Put(user, module, key, right)
	rightresult, err := cache.Get(user, module, key, right)
	assert.EqualValues(t, right, rightresult)
	assert.Nil(t, err)
	cache.Invalidate(user)
	ok, err := cache.Get(user, module, key, right)
	assert.EqualValues(t, ok, false)
	assert.NotNil(t, err)

}
