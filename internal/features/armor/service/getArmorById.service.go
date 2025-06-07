package armor_service

import (
	"context"
	"log"

	armor_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *ArmorService) GetArmorById(ctx context.Context, id string) (*model.Armor, error) {
	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Println("ArmorService.GetArmorById Error : ", err)
		return nil, err
	}

	result, err := s.armorRepository.FindOneById(ctx, objId)
	if err != nil {
		log.Println("ArmorService.GetArmorById Error : ", err)
		return nil, err
	}

	afterMap := armor_mapper.MapperArmorEntityToModel(*result)
	return &afterMap, nil
}
