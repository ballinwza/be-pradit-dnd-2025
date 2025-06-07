package core_mapper

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperWeightEntityToModel(entity core_outbound_entity.WeightEntity) model.Weight {
	return model.Weight{
		Unit:  entity.Unit,
		Value: entity.Value,
	}
}
