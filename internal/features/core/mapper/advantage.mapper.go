package core_mapper

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperAdvantageTypeEntityToModel(entity core_outbound_entity.AdvantageTypeEntity) model.AdvantageType {
	switch entity {
	case core_outbound_entity.AdvantageTypeDisadvantage:
		return model.AdvantageTypeDisadvantage
	case core_outbound_entity.AdvantageTypeNone:
		return model.AdvantageTypeNone
	case core_outbound_entity.AdvantageTypeAdvantage:
		return model.AdvantageTypeAdvantage
	default:
		return "unknow AdvantageType"
	}
}
