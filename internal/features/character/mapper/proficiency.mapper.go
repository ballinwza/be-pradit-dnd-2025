package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperProficiencyEntityToModel(entity character_outbound_entity.ProficiencyEntity) model.CharacterProficiency {

	return model.CharacterProficiency{
		Name:                  core_mapper.MapperProficiencyTypeEntityToModel(entity.Name),
		Value:                 entity.Value,
		AbilityShortTypeGroup: core_mapper.MapperAbilityShortTypeEntityToModel(entity.AbilityShortGroup),
	}

}
