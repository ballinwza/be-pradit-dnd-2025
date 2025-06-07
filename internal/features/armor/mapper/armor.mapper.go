package armor_mapper

import (
	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperArmorEntityToModel(entity armor_outbound_entity.ArmorEntity) model.Armor {
	afterMapCoin := core_mapper.MapperCoinEntityToModel(entity.Price)
	afterMapWeight := core_mapper.MapperWeightEntityToModel(entity.Weight)

	if entity.MaximumPlusArmorClass == nil {
		defaultValue := int32(10)
		entity.MaximumPlusArmorClass = &defaultValue
	}

	return model.Armor{
		ID:                    entity.Id.Hex(),
		ArmorType:             MapperArmorTypeEntityToModel(entity.ArmorType),
		Name:                  entity.Name,
		ArmorClass:            entity.ArmorClass,
		MaximumPlusArmorClass: entity.MaximumPlusArmorClass,
		StealthAdvantageType:  core_mapper.MapperAdvantageTypeEntityToModel(entity.StealthAdvantage),
		Price:                 &afterMapCoin,
		StrRequirement:        entity.StrRequirement,
		DescriptionEn:         entity.DescriptionEn,
		Weight:                &afterMapWeight,
	}
}
