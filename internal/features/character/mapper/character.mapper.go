package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterEntityToModel(entity character_outbound_entity.CharacterEntity) model.Character {
	afterMappingHitPoint := MapperHitPointEntityToModel(entity.HitPoint)

	var afterMappingProficiency []*model.CharacterProficiency
	var afterMappingAbility []*model.CharacterAbility
	var afterMappingPocketMoney []*model.Coin

	for _, value := range entity.Ability {
		var abilityModel model.CharacterAbility = MapperAbilityEntityToModel(*value)
		afterMappingAbility = append(afterMappingAbility, &abilityModel)
	}
	for _, value := range entity.PocketMoney {
		coinModel := core_mapper.MapperCoinEntityToModel(*value)
		afterMappingPocketMoney = append(afterMappingPocketMoney, coinModel)
	}
	for _, value := range entity.Proficiency {
		var proficiencyModel model.CharacterProficiency = MapperCharacterProficiencyEntityToModel(*value)
		afterMappingProficiency = append(afterMappingProficiency, &proficiencyModel)
	}

	chkId := entity.Id.Hex()
	return model.Character{
		ID:          chkId,
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

func MapperCharacterModelToEntity(req model.SaveCharacterReq) (*character_outbound_entity.CharacterEntity, error) {
	afterMappingHitPoint := MapperHitPointModelToEntity(*req.HitPoint)

	var afterMappingProficiency []*character_outbound_entity.CharacterProficiencyEntity
	var afterMappingAbility []*character_outbound_entity.AbilityEntity
	var afterMappingPocketMoney []*core_outbound_entity.CoinEntity

	for _, value := range req.Ability {
		abilityModel := MapperAbilityModelToEntity(*value)
		afterMappingAbility = append(afterMappingAbility, abilityModel)
	}
	for _, value := range req.PocketMoney {
		coinModel := core_mapper.MapperCoinModelToEntity(*value)
		afterMappingPocketMoney = append(afterMappingPocketMoney, coinModel)
	}
	for _, value := range req.Proficiency {
		proficiencyModel := MapperCharacterProficiencyModelToEntity(*value)
		afterMappingProficiency = append(afterMappingProficiency, proficiencyModel)
	}

	var objectId *bson.ObjectID
	if req.ID != nil {
		changedIdType, _ := bson.ObjectIDFromHex(*req.ID)
		objectId = &changedIdType
	}

	var classObjId *bson.ObjectID
	if req.ClassID != nil {
		objId, _ := bson.ObjectIDFromHex(*req.ClassID)
		classObjId = &objId
	}

	result := character_outbound_entity.CharacterEntity{
		Id:          objectId,
		Name:        req.Name,
		HitPoint:    afterMappingHitPoint,
		CurrentExp:  *req.CurrentExp,
		AvatarImage: *req.AvatarImage,
		PocketMoney: afterMappingPocketMoney,
		Proficiency: afterMappingProficiency,
		Ability:     afterMappingAbility,
		ClassId:     *classObjId,
	}

	return &result, nil
}
