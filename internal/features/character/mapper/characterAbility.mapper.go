package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperAbilityEntityToModel(entity character_outbound_entity.AbilityEntity) model.CharacterAbility {

	return model.CharacterAbility{
		Name:      entity.Name,
		Value:     *entity.Value,
		ShortType: core_mapper.MapperAbilityShortTypeEntityToModel(entity.ShortType),
	}
}
func MapperAbilityModelToEntity(model model.CharacterAbilityReq) *character_outbound_entity.AbilityEntity {
	name, _ := core_mapper.MapperAbilityShortTypeToFullnameStringModelToEntity(model.ShortType)
	return &character_outbound_entity.AbilityEntity{
		Name:      name,
		Value:     &model.Value,
		ShortType: core_mapper.MapperAbilityShortTypeModelToEntity(model.ShortType),
	}
}
