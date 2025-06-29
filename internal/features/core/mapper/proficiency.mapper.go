package core_mapper

import (
	"fmt"

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
func MapperProficiencyTypeModelToEntity(proficiencyModel model.ProficiencyType) core_outbound_entity.ProficiencyType {
	switch proficiencyModel {
	case model.ProficiencyTypeAthletics:
		return core_outbound_entity.ProficiencyTypeAthletics
	case model.ProficiencyTypeArobatics:
		return core_outbound_entity.ProficiencyTypeArobatic
	case model.ProficiencyTypeSleightofhand:
		return core_outbound_entity.ProficiencyTypeSleightOfHand
	case model.ProficiencyTypeStealth:
		return core_outbound_entity.ProficiencyTypeStealth
	case model.ProficiencyTypeArcana:
		return core_outbound_entity.ProficiencyTypeArcana
	case model.ProficiencyTypeHistory:
		return core_outbound_entity.ProficiencyTypeHistory
	case model.ProficiencyTypeInvestigation:
		return core_outbound_entity.ProficiencyTypeInvestigation
	case model.ProficiencyTypeNature:
		return core_outbound_entity.ProficiencyTypeNature
	case model.ProficiencyTypeReligion:
		return core_outbound_entity.ProficiencyTypeReligion
	case model.ProficiencyTypeAnimalhandling:
		return core_outbound_entity.ProficiencyTypeAnimalHandling
	case model.ProficiencyTypeInsight:
		return core_outbound_entity.ProficiencyTypeInsight
	case model.ProficiencyTypeMedicine:
		return core_outbound_entity.ProficiencyTypeMedicine
	case model.ProficiencyTypePerception:
		return core_outbound_entity.ProficiencyTypePerception
	case model.ProficiencyTypeSurvival:
		return core_outbound_entity.ProficiencyTypeSurvival
	case model.ProficiencyTypeDeception:
		return core_outbound_entity.ProficiencyTypeDeception
	case model.ProficiencyTypeIntimidation:
		return core_outbound_entity.ProficiencyTypeIntimidation
	case model.ProficiencyTypePerformance:
		return core_outbound_entity.ProficiencyTypePerformance
	case model.ProficiencyTypePersuasion:
		return core_outbound_entity.ProficiencyTypePersuasion
	default:
		return "unknow abilityShortType"
	}
}

func MapperProficiencyTypeModelToEntityWithFullname(proficiencyModel model.ProficiencyType) string {
	switch proficiencyModel {
	case model.ProficiencyTypeAthletics:
		return "athletics"
	case model.ProficiencyTypeArobatics:
		return "arobatics"
	case model.ProficiencyTypeSleightofhand:
		return "sleight_of_hand"
	case model.ProficiencyTypeStealth:
		return "stealth"
	case model.ProficiencyTypeArcana:
		return "arcana"
	case model.ProficiencyTypeHistory:
		return "history"
	case model.ProficiencyTypeInvestigation:
		return "investigation"
	case model.ProficiencyTypeNature:
		return "nature"
	case model.ProficiencyTypeReligion:
		return "religion"
	case model.ProficiencyTypeAnimalhandling:
		return "animal_handling"
	case model.ProficiencyTypeInsight:
		return "insight"
	case model.ProficiencyTypeMedicine:
		return "medicine"
	case model.ProficiencyTypePerception:
		return "perception"
	case model.ProficiencyTypeSurvival:
		return "survival"
	case model.ProficiencyTypeDeception:
		return "deception"
	case model.ProficiencyTypeIntimidation:
		return "initimidation"
	case model.ProficiencyTypePerformance:
		return "performance"
	case model.ProficiencyTypePersuasion:
		return "persuasion"
	default:
		fmt.Printf("Error MapperProficiencyTypeModelToEntityWithFullname not match any type : %v", proficiencyModel)
		return ""
	}
}
