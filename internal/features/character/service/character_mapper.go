package character_service

import (
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/entity"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/service"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperEntityToModel(entity *character_entity.CharacterEntity) *model.Character {
	return &model.Character{
		ID: entity.Id.Hex(),
		User: user_service.MapperEntityToModel(&entity.UserId),
		Name: entity.Name,
		AvatarImage: entity.AvatarImage,
		Speed: entity.Speed,
		InitiativePoint: entity.InitiativePlusPoint,
		HitDice: entity.HitDice,
	}
}