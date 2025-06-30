package character_outbound_repository

import (
	"context"
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
)

func (r CharacterRepository) InsertOne(ctx context.Context, characterEntity character_outbound_entity.CharacterEntity) (bool, error) {

	factory := character_outbound_entity.CharacterEntity{
		Name:        characterEntity.Name,
		HitPoint:    characterEntity.HitPoint,
		CurrentExp:  characterEntity.CurrentExp,
		AvatarImage: characterEntity.AvatarImage,
		PocketMoney: characterEntity.PocketMoney,
		ClassId:     characterEntity.ClassId,
		Ability:     characterEntity.Ability,
		Proficiency: characterEntity.Proficiency,
	}

	result, err := r.collection.InsertOne(ctx, factory)
	if err != nil {
		log.Println("Error CharacterRepository.InsertOne : ", err)
		return false, err
	}

	return result.Acknowledged, nil
}
