package core_mapper

import (
	"fmt"

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

func MapperAbilityShortTypeModelToEntity(typeFromModel model.AbilityShortType) core_outbound_entity.AbilityType {
	switch typeFromModel {
	case model.AbilityShortTypeStr:
		return core_outbound_entity.AbilityTypeStr
	case model.AbilityShortTypeDex:
		return core_outbound_entity.AbilityTypeDex
	case model.AbilityShortTypeCon:
		return core_outbound_entity.AbilityTypeCon
	case model.AbilityShortTypeInt:
		return core_outbound_entity.AbilityTypeInt
	case model.AbilityShortTypeWis:
		return core_outbound_entity.AbilityTypeWis
	case model.AbilityShortTypeCha:
		return core_outbound_entity.AbilityTypeCha
	default:
		return "unknow abilityShortType"
	}
}

func MapperAbilityShortTypeToFullnameStringModelToEntity(typeFromModel model.AbilityShortType) (string, error) {
	switch typeFromModel {
	case model.AbilityShortTypeStr:
		return "strength", nil
	case model.AbilityShortTypeDex:
		return "dexterity", nil
	case model.AbilityShortTypeCon:
		return "constitution", nil
	case model.AbilityShortTypeInt:
		return "intelligence", nil
	case model.AbilityShortTypeWis:
		return "wisdom", nil
	case model.AbilityShortTypeCha:
		return "charisma", nil
	default:
		return "", fmt.Errorf("unknown ability short type: %v", typeFromModel)
	}
}
