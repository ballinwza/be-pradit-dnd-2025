package ability_detail_mapper

import (
	ability_detail_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperAbilityDetailEntityToModel(entity ability_detail_outbound_entity.AbilityDetailEntity) model.AbilityDetail {
	proficiencyModelList := make([]*model.ProficiencyDetail, len(entity.Proficiencies))

	for i, proficiency := range entity.Proficiencies {
		afterMap := MapperProficiencyDetailEntityToModel(proficiency)
		proficiencyModelList[i] = &afterMap
	}

	return model.AbilityDetail{
		ID:            entity.Id.Hex(),
		Name:          entity.Name,
		Short:         MapperShortTypeEntityToModel(entity.Short),
		DescriptionEn: entity.DescriptionEn,
		DescriptionTh: entity.DescriptionTh,
		Proficiencies: proficiencyModelList,
	}
}

func MapperAbilityDetailShortTypeModelToEntity(short model.AbilityShortType) ability_detail_outbound_entity.AbilityDetailShortTypeEntity {
	switch short {
	case model.AbilityShortTypeStr:
		return ability_detail_outbound_entity.AbilityDetailShortTypeEntityStr
	case model.AbilityShortTypeDex:
		return ability_detail_outbound_entity.AbilityDetailShortTypeEntityDex
	case model.AbilityShortTypeCon:
		return ability_detail_outbound_entity.AbilityDetailShortTypeEntityCon
	case model.AbilityShortTypeInt:
		return ability_detail_outbound_entity.AbilityDetailShortTypeEntityInt
	case model.AbilityShortTypeWis:
		return ability_detail_outbound_entity.AbilityDetailShortTypeEntityWis
	case model.AbilityShortTypeCha:
		return ability_detail_outbound_entity.AbilityDetailShortTypeEntityCha
	default:
		return "unknow GraphQL shortType "
	}
}

func MapperShortTypeEntityToModel(short ability_detail_outbound_entity.AbilityDetailShortTypeEntity) model.AbilityShortType {
	switch short {
	case ability_detail_outbound_entity.AbilityDetailShortTypeEntityStr:
		return model.AbilityShortTypeStr
	case ability_detail_outbound_entity.AbilityDetailShortTypeEntityDex:
		return model.AbilityShortTypeDex
	case ability_detail_outbound_entity.AbilityDetailShortTypeEntityCon:
		return model.AbilityShortTypeCon
	case ability_detail_outbound_entity.AbilityDetailShortTypeEntityInt:
		return model.AbilityShortTypeInt
	case ability_detail_outbound_entity.AbilityDetailShortTypeEntityWis:
		return model.AbilityShortTypeWis
	case ability_detail_outbound_entity.AbilityDetailShortTypeEntityCha:
		return model.AbilityShortTypeCha
	default:
		return "unknow GraphQL shortType "
	}
}
