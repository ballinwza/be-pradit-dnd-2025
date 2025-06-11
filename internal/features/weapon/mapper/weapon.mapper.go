package weapon_mapper

import (
	"log"

	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"
	weapon_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperWeaponEntityToModel(entity weapon_outbound_entity.WeaponEntity) model.Weapon {
	afterMapperWeight := core_mapper.MapperWeightEntityToModel(entity.Weight)
	afterMappingPrice := core_mapper.MapperCoinEntityToModel(entity.Price)
	afterMappingMastery := MapperMasteryEntityToModel(entity.MasteryId)

	afterMappingDiceRoll, err := core_mapper.MapperDiceRollTypeEntityToModel(entity.DiceRollType)
	if err != nil {
		log.Println("MapperWeaponEntityToModel.MapperDiceRollTypeEntityToModel Error : ", err)
	}

	afterMappingDamageType, err := core_mapper.MapperDamagedTypeEntityToModel(entity.DamagesType)
	if err != nil {
		log.Println("MapperWeaponEntityToModel.MapperDamagedTypeEntityToModel Error : ", err)
	}

	propertyListAfterMapping := make([]*model.WeaponProperty, len(entity.PropertyId))
	for i, weapon := range entity.PropertyId {
		afterMapping := MapperPropertyEntityToModel(weapon)
		propertyListAfterMapping[i] = &afterMapping
	}

	return model.Weapon{
		ID:             entity.Id.Hex(),
		Name:           entity.Name,
		DiceQuantity:   entity.DiceQuantity,
		DescriptionEn:  entity.DescriptionEn,
		DescriptionTh:  entity.DescriptionTh,
		DiceRollType:   afterMappingDiceRoll,
		DamagedType:    afterMappingDamageType,
		Weight:         &afterMapperWeight,
		Price:          &afterMappingPrice,
		WeaponMastery:  &afterMappingMastery,
		WeaponProperty: propertyListAfterMapping,
		ImageURL:       entity.ImageUrl,
		NormalRange:    &entity.NormalRange,
		LongRange:      &entity.LongRange,
	}
}
