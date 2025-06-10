package core_mapper

import (
	"fmt"

	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperDiceRollTypeEntityToModel(entity core_outbound_entity.DiceRollTypeEntity) (model.DiceRollType, error) {
	switch entity {
	case core_outbound_entity.DiceRollTypeEntityD20:
		return model.DiceRollTypeD20, nil
	case core_outbound_entity.DiceRollTypeEntityD12:
		return model.DiceRollTypeD12, nil
	case core_outbound_entity.DiceRollTypeEntityD10:
		return model.DiceRollTypeD10, nil
	case core_outbound_entity.DiceRollTypeEntityD8:
		return model.DiceRollTypeD8, nil
	case core_outbound_entity.DiceRollTypeEntityD6:
		return model.DiceRollTypeD6, nil
	case core_outbound_entity.DiceRollTypeEntityD4:
		return model.DiceRollTypeD4, nil
	default:
		var zeroValue model.DiceRollType
		return zeroValue, fmt.Errorf("unsupported or unknown DiceRollTypeEntity: %v", entity)
	}
}
