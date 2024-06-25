package cacher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	tests := []struct {
		name           string
		actions        func(CacheInterface)
		expectedResult bool
		expectedError  bool
	}{
		{
			name: "Put and Get cache entry",
			actions: func(c CacheInterface) {
				c.Put("myrachanto", "users", "READ", true)
			},
			expectedResult: true,
			expectedError:  false,
		},
		{
			name: "Invalidate cache entry and Get",
			actions: func(c CacheInterface) {
				c.Put("myrachanto", "users", "READ", true)
				c.Invalidate("myrachanto")
			},
			expectedResult: false,
			expectedError:  true,
		},
		{
			name: "Get non-existing user",
			actions: func(c CacheInterface) {
			},
			expectedResult: false,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache()
			tt.actions(cache)
			result, err := cache.Get("myrachanto", "users", "READ")

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
