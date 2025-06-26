package character_repository

import (
	"context"
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *CharacterRepository) UpdateOneById(ctx context.Context, id bson.ObjectID, updateValue character_outbound_entity.CharacterEntity) (bool, error) {
	result, err := r.collection.UpdateByID(ctx, bson.M{"_id": id}, bson.M{"$set": updateValue})
	if err != nil {
		log.Println("Error CharacterRepository.UpdateOneById : ", err)
		return false, err
	}

	return result.ModifiedCount > 0, nil
}
