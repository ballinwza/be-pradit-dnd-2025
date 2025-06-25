package core_mapper

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperProficiencyTypeEntityToModel(entity core_outbound_entity.ProficiencyType) model.ProficiencyType {
	switch entity {
	case core_outbound_entity.ProficiencyTypeAthletics:
		return model.ProficiencyTypeAthletics
	case core_outbound_entity.ProficiencyTypeArobatic:
		return model.ProficiencyTypeArobatics
	case core_outbound_entity.ProficiencyTypeSleightOfHand:
		return model.ProficiencyTypeSleightofhand
	case core_outbound_entity.ProficiencyTypeStealth:
		return model.ProficiencyTypeStealth
	case core_outbound_entity.ProficiencyTypeArcana:
		return model.ProficiencyTypeArcana
	case core_outbound_entity.ProficiencyTypeHistory:
		return model.ProficiencyTypeHistory
	case core_outbound_entity.ProficiencyTypeInvestigation:
		return model.ProficiencyTypeInvestigation
	case core_outbound_entity.ProficiencyTypeNature:
		return model.ProficiencyTypeNature
	case core_outbound_entity.ProficiencyTypeReligion:
		return model.ProficiencyTypeReligion
	case core_outbound_entity.ProficiencyTypeAnimalHandling:
		return model.ProficiencyTypeAnimalhandling
	case core_outbound_entity.ProficiencyTypeInsight:
		return model.ProficiencyTypeInsight
	case core_outbound_entity.ProficiencyTypeMedicine:
		return model.ProficiencyTypeMedicine
	case core_outbound_entity.ProficiencyTypePerception:
		return model.ProficiencyTypePerception
	case core_outbound_entity.ProficiencyTypeSurvival:
		return model.ProficiencyTypeSurvival
	case core_outbound_entity.ProficiencyTypeDeception:
		return model.ProficiencyTypeDeception
	case core_outbound_entity.ProficiencyTypeIntimidation:
		return model.ProficiencyTypeIntimidation
	case core_outbound_entity.ProficiencyTypePerformance:
		return model.ProficiencyTypePerformance
	case core_outbound_entity.ProficiencyTypePersuasion:
		return model.ProficiencyTypePersuasion
	default:
		return "unknow abilityShortType"
	}
}
