package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterProficiencyEntityToModel(entity character_outbound_entity.CharacterProficiencyEntity) model.CharacterProficiency {
	return model.CharacterProficiency{
		Name:                  core_mapper.MapperProficiencyTypeEntityToModel(entity.Name),
		Value:                 *entity.Value,
		AbilityShortTypeGroup: core_mapper.MapperAbilityShortTypeEntityToModel(entity.AbilityShortGroup),
	}
}
func MapperCharacterProficiencyModelToEntity(proficiencyModel model.CharacterProficiencyReq) *character_outbound_entity.CharacterProficiencyEntity {
	return &character_outbound_entity.CharacterProficiencyEntity{
		Name:              core_mapper.MapperProficiencyTypeModelToEntity(proficiencyModel.Name),
		Value:             &proficiencyModel.Value,
		AbilityShortGroup: core_mapper.MapperAbilityShortTypeModelToEntity(proficiencyModel.AbilityShortTypeGroup),
	}
}
