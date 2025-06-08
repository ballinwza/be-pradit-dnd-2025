package armor_service

import (
	"context"
	"log"

	armor_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/mapper"
	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *ArmorService) GetShieldList(ctx context.Context) ([]*model.Armor, error) {
	shieldList, err := s.armorRepository.FindAll(ctx, armor_outbound_entity.ArmorEquipmentTypeShield)
	if err != nil {
		log.Println("ArmorService.GetShieldList Error : ", err)
		return nil, err
	}

	result := make([]*model.Armor, len(shieldList))
	for i, armor := range shieldList {
		afterMapping := armor_mapper.MapperArmorEntityToModel(*armor)
		result[i] = &afterMapping
	}

	return result, nil
}
