package character_service

import (
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/character/entity"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/user/service"
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