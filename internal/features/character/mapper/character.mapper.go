package character_mapper

import (
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterEntityToModel(entity character_entity.CharacterEntity) model.Character {
	afterMappingHitPoint := MapperHitPointEntityToModel(entity.HitPoint)
	afterMappingProficiency := MapperProficiencyEntityToModel(entity.Proficiency)
	afterMappingAbility := MapperAbilityEntityToModel(entity.Ability)
	var afterMappingPocketMoney []*model.Coin

	for _, value := range entity.PocketMoney {
		var coinModel model.Coin = core_mapper.MapperCoinEntityToModel(value)
		afterMappingPocketMoney = append(afterMappingPocketMoney, &coinModel)
	}

	return model.Character{
		ID:          entity.Id.Hex(),
		Name:        entity.Name,
		HitPoint:    &afterMappingHitPoint,
		CurrentExp:  entity.CurrentExp,
		AvatarImage: entity.AvatarImage,
		PocketMoney: afterMappingPocketMoney,
		Proficiency: &afterMappingProficiency,
		Ability:     &afterMappingAbility,
	}
}
