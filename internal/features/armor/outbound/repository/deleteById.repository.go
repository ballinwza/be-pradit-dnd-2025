package armor_outbound_repository

import (
	"context"
	"log"

	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *ArmorRepository) DeleteById(ctx context.Context, objId bson.ObjectID) (bool, error) {
	filter := armor_outbound_entity.ArmorEntityFilter{
		Id: &objId,
	}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error ArmorRepository.DeleteById can't delete ID : %v , %v", objId, err)
		return false, err
	}

	return result.Acknowledged, nil
}
