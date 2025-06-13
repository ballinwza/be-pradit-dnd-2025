package equipment_service

import (
	"context"
	"log"

	equipment_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *EquipmentService) GetEquipmentByCharacterId(ctx context.Context, characterId string) (*model.Equipment, error) {
	objId, err := bson.ObjectIDFromHex(characterId)
	if err != nil {
		log.Printf("EquipmentService.GetEquipmentByCharacterId Invalid ObjectId : %v\n", err)
		return nil, err
	}

	result, err := s.equipmentRepository.FindByCharacterId(ctx, objId)
	if err != nil {
		log.Printf("EquipmentService.GetEquipmentByCharacterId Error : %v\n", err)
		return nil, err
	}

	resultMapping := equipment_mapper.MapperEquipmentEntityToModel(*result)
	return &resultMapping, nil
}
