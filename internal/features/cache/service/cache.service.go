package cache_service

import "github.com/patrickmn/go-cache"

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService(cache *cache.Cache) *CacheService {
	return &CacheService{
		cache: cache,
	}
}
