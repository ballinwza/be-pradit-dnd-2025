package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterEntityToModel(entity character_outbound_entity.CharacterEntity) model.Character {
	afterMappingHitPoint := MapperHitPointEntityToModel(entity.HitPoint)

	var afterMappingProficiency []*model.CharacterProficiency
	var afterMappingAbility []*model.CharacterAbility
	var afterMappingPocketMoney []*model.Coin

	for _, value := range entity.Ability {
		var abilityModel model.CharacterAbility = MapperAbilityEntityToModel(value)
		afterMappingAbility = append(afterMappingAbility, &abilityModel)
	}
	for _, value := range entity.PocketMoney {
		var coinModel model.Coin = core_mapper.MapperCoinEntityToModel(value)
		afterMappingPocketMoney = append(afterMappingPocketMoney, &coinModel)
	}
	for _, value := range entity.Proficiency {
		var proficiencyModel model.CharacterProficiency = MapperProficiencyEntityToModel(value)
		afterMappingProficiency = append(afterMappingProficiency, &proficiencyModel)
	}
	return model.Character{
		ID:          entity.Id.Hex(),
		Name:        entity.Name,
		HitPoint:    &afterMappingHitPoint,
		CurrentExp:  entity.CurrentExp,
		AvatarImage: entity.AvatarImage,
		PocketMoney: afterMappingPocketMoney,
		Proficiency: afterMappingProficiency,
		Ability:     afterMappingAbility,
		ClassID:     entity.ClassId.Hex(),
	}
}
