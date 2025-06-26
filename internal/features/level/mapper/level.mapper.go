package level_mapper

import (
	level_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/level/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperLevelEntityToModel(entity level_outbound_entity.LevelEntity) model.Level {
	convertIdToString := entity.Id.Hex()

	return model.Level{
		ID:                    &convertIdToString,
		Level:                 entity.Level,
		ExpToLevelUp:          entity.ExpToLevelUp,
		ProficiencyBonusPoint: entity.ProficiencyBonusPoint,
	}
}
