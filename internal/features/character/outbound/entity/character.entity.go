package character_outbound_entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CharacterEntity struct {
	Id          *bson.ObjectID             `bson:"_id,omitempty"`
	Name        string                     `bson:"name"`
	HitPoint    HitPointEntity             `bson:"hit_point"`
	CurrentExp  int32                      `bson:"current_exp"`
	AvatarImage string                     `bson:"avatar_image"`
	PocketMoney CharacterPocketMoneyEntity `bson:"pocket_money"`
	Proficiency CharacterProficiencyEntity `bson:"proficiencies"`
	Ability     AbilityEntity              `bson:"abilities"`
	ClassId     bson.ObjectID              `bson:"class_id"`
}
