package armor_service

import (
	armor_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/repository"
	"github.com/patrickmn/go-cache"
)

type ArmorService struct {
	armorRepository *armor_outbound_repository.ArmorRepository
	cache           *cache.Cache
	// redisClient     *redis.Client
}

func NewArmorService(cache *cache.Cache) *ArmorService {
	armorRepository := armor_outbound_repository.NewArmorRepository()
	// redisClient, err := cache_service.NewRedisClient()
	// if err != nil {
	// 	log.Printf("Error Impl NewArmorRepository can't connect cache : %v", err)
	// }
	// inMemoryCache := cache_service.NewInMemoryCache(20*time.Second, 30*time.Second)

	return &ArmorService{
		armorRepository: armorRepository,
		// redisClient:     redisClient,
		cache: cache,
	}
}
