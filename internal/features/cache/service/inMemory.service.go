package cache_service

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func NewInMemoryCache(defaultExpiration, cleanupInterval time.Duration) *cache.Cache {
	c := cache.New(defaultExpiration, cleanupInterval)
	return c
}
