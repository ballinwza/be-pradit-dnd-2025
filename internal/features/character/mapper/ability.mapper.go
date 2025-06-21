package character_mapper

import (
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperAbilityEntityToModel(entity character_entity.AbilityEntity) model.Ability {

	return model.Ability{
		Strength:     entity.Strength,
		Dexterity:    entity.Dexterity,
		Constitution: entity.Constitution,
		Intelligence: entity.Intelligence,
		Wisdom:       entity.Wisdom,
		Charisma:     entity.Charisma,
	}
}
