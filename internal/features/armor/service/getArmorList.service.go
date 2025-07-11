package armor_service

import (
	"context"
	"log"

	armor_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/mapper"
	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"github.com/patrickmn/go-cache"
)

func (s *ArmorService) GetArmorList(ctx context.Context) ([]*model.Armor, error) {
	cacheKey := "armor_list_all"

	if data, found := s.cache.Get(cacheKey); found {
		log.Println("✅ In-Memory Cache Hit for key:", cacheKey)
		if armorList, ok := data.([]*model.Armor); ok {
			return armorList, nil
		}
	}

	log.Println("❌ In-Memory Cache Miss for key:", cacheKey)

	armorList, err := s.armorRepository.FindAll(ctx, armor_outbound_entity.ArmorEquipmentTypeArmor)
	if err != nil {
		log.Println("Error ArmorService.GetArmorList FindAll : ", err)
		return nil, err
	}

	result := make([]*model.Armor, len(armorList))
	for i, armor := range armorList {
		afterMapping := armor_mapper.MapperArmorEntityToModel(*armor)
		result[i] = &afterMapping
	}

	log.Println(" caching data for key:", cacheKey)
	s.cache.Set(cacheKey, result, cache.DefaultExpiration)

	return result, nil
}
