package character_outbound_repository

import (
	"context"
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *CharacterRepository) FindAllByUserId(ctx context.Context, userId bson.ObjectID) ([]*character_outbound_entity.CharacterEntity, error) {
	filter := character_outbound_entity.CharacterEntityFilter{
		UserId: &userId,
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Println("Error CharacterRepository.FindAllByUserId not founded character : ", err)
		return nil, err
	}

	defer cursor.Close(ctx)
	var result []*character_outbound_entity.CharacterEntity

	if err = cursor.All(ctx, &result); err != nil {
		log.Println("Error CharacterRepository.FindAllByUserId Decoding : ", err)
		return nil, err
	}

	return result, nil
}
