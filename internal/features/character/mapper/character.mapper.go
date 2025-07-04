package character_mapper

import (
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterEntityToModel(entity character_outbound_entity.CharacterEntity) model.Character {
	afterMappingHitPoint := MapperHitPointEntityToModel(entity.HitPoint)
	afterMappingPocketMoney := MapperCharacterPocketMoneyEntityToModel(entity.PocketMoney)

	chkId := entity.Id.Hex()
	return model.Character{
		ID:          chkId,
		Name:        entity.Name,
		HitPoint:    &afterMappingHitPoint,
		CurrentExp:  entity.CurrentExp,
		AvatarImage: entity.AvatarImage,
		PocketMoney: afterMappingPocketMoney,
		Proficiency: (*model.CharacterProficiency)(&entity.Proficiency),
		Ability:     (*model.CharacterAbility)(&entity.Ability),
		ClassID:     entity.ClassId.Hex(),
		UserID:      entity.UserId.Hex(),
	}
}

func MapperCharacterModelToEntity(req model.SaveCharacterReq) (*character_outbound_entity.CharacterEntity, error) {
	afterMappingHitPoint := MapperHitPointModelToEntity(*req.HitPoint)
	afterMapAbility := MapperAbilityModelToEntity(req.Ability)
	afterMapProficiency := MapperCharacterProficiencyModelToEntity(*req.Proficiency)
	afterMappingPocketMoney := MapperCharacterPocketMoneyModelToEntity(*req.PocketMoney)

	var objectId *bson.ObjectID
	if req.ID != nil {
		changedIdType, err := bson.ObjectIDFromHex(*req.ID)
		objectId = &changedIdType
		log.Printf("Error MapperCharacterModelToEntity invalid objecId : %v", err)
	}

	classObjId, _ := bson.ObjectIDFromHex(req.ClassID)
	userIdObjId, _ := bson.ObjectIDFromHex(req.UserID)

	result := character_outbound_entity.CharacterEntity{
		Id:          objectId,
		Name:        req.Name,
		HitPoint:    afterMappingHitPoint,
		CurrentExp:  req.CurrentExp,
		AvatarImage: req.AvatarImage,
		PocketMoney: *afterMappingPocketMoney,
		Proficiency: *afterMapProficiency,
		Ability:     *afterMapAbility,
		ClassId:     classObjId,
		UserId:      userIdObjId,
	}

	return &result, nil
}
