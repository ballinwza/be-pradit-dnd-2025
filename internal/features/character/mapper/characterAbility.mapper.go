package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperAbilityEntityToModel(entity character_outbound_entity.AbilityEntity) model.CharacterAbility {
	return model.CharacterAbility{
		Str: entity.Str,
		Dex: entity.Dex,
		Con: entity.Con,
		Int: entity.Int,
		Wis: entity.Wis,
		Cha: entity.Cha,
	}
}
func MapperAbilityModelToEntity(model *model.CharacterAbilityReq) *character_outbound_entity.AbilityEntity {
	return &character_outbound_entity.AbilityEntity{
		Str: model.Str,
		Dex: model.Dex,
		Con: model.Con,
		Int: model.Int,
		Wis: model.Wis,
		Cha: model.Cha,
	}
}
