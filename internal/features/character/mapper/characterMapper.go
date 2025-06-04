package character_mapper

import (
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/service"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterEntityToModel(entity character_entity.CharacterEntity) model.Character {
	return model.Character{
		ID: entity.Id.Hex(),
		User: user_service.MapperEntityToModel(&entity.UserId),
		Name: entity.Name,
		AvatarImage: entity.AvatarImage,
		Speed: entity.Speed,
		InitiativePoint: entity.InitiativePlusPoint,
		HitDice: entity.HitDice,
	}
}