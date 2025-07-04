package character_outbound_repository

import (
	"context"
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
)

func (r CharacterRepository) InsertOne(ctx context.Context, characterEntity character_outbound_entity.CharacterEntity) (bool, error) {
	if err := r.validate.Struct(characterEntity); err != nil {
		log.Println("Error CharacterRepository.InsertOne validate: ", err)
		return false, err
	}

	result, err := r.collection.InsertOne(ctx, characterEntity)
	if err != nil {
		log.Println("Error CharacterRepository.InsertOne can't insert character : ", err)
		return false, err
	}

	return result.Acknowledged, nil
}
