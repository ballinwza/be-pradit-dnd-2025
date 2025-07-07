package cache_service

func (s *CacheService) ClearInMemoryCache() {
	s.cache.Flush()
}
