package armor_service

import (
	"context"

	armor_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *ArmorService) GetArmorList(ctx context.Context) ([]*model.Armor, error) {
	armorList, err := s.armorRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Armor, len(armorList))
	for i, armor := range armorList {
		afterMapping := armor_mapper.MapperArmorEntityToModel(*armor)
		result[i] = &afterMapping
	}

	return result, nil
}
