package core_mapper

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperAbilityShortTypeEntityToModel(entity core_outbound_entity.AbilityType) model.AbilityShortType {
	switch entity {
	case core_outbound_entity.AbilityTypeStr:
		return model.AbilityShortTypeStr
	case core_outbound_entity.AbilityTypeDex:
		return model.AbilityShortTypeDex
	case core_outbound_entity.AbilityTypeCon:
		return model.AbilityShortTypeCon
	case core_outbound_entity.AbilityTypeInt:
		return model.AbilityShortTypeInt
	case core_outbound_entity.AbilityTypeWis:
		return model.AbilityShortTypeWis
	case core_outbound_entity.AbilityTypeCha:
		return model.AbilityShortTypeCha
	default:
		return "unknow abilityShortType"
	}
}
