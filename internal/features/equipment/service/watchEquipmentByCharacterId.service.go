package equipment_service

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *EquipmentService) WatchEquipmentByCharacterId(ctx context.Context, characterId string) (<-chan *model.Equipment, error) {
	objId, err := bson.ObjectIDFromHex(characterId)
	if err != nil {
		log.Printf("EquipmentService.GetEquipmentByCharacterId Invalid ObjectId : %v\n", err)
		return nil, err
	}

	repoChan, err := s.equipmentRepository.WatchByCharacterId(ctx, objId)
	if err != nil {
		log.Printf("EquipmentService.GetEquipmentByCharacterId Error : %v\n", err)
		return nil, err
	}

	return repoChan, nil
}
