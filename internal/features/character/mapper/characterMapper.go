package character_mapper

import (
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	user_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/mapper"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterEntityToModel(entity character_entity.CharacterEntity) model.Character {
	userAfterMapping := user_mapper.MapperUserEntityToModel(entity.UserId)
	
	return model.Character{
		ID: entity.Id.Hex(),
		User: &userAfterMapping,
		Name: entity.Name,
		AvatarImage: entity.AvatarImage,
		Speed: entity.Speed,
		InitiativePoint: entity.InitiativePlusPoint,
		HitDice: entity.HitDice,
	}
}