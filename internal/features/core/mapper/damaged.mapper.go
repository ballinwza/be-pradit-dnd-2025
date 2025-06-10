package core_mapper

import (
	"fmt"

	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperDamagedTypeEntityToModel(entity core_outbound_entity.DamagedType) (model.DamagedType, error) {
	switch entity {
	case core_outbound_entity.DamagedTypeBludgeon:
		return model.DamagedTypeBludgeon, nil
	case core_outbound_entity.DamagedTypePierce:
		return model.DamagedTypePierce, nil
	case core_outbound_entity.DamagedTypeSlash:
		return model.DamagedTypeSLASh, nil
	default:
		var zeroValue model.DamagedType
		return zeroValue, fmt.Errorf("unsupported or unknown DamagedTypeEntity: %v", entity)
	}
}
