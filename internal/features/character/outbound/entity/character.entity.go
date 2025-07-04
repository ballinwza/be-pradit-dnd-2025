package character_outbound_entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CharacterEntity struct {
	Id          *bson.ObjectID             `validate:"omitempty" bson:"_id,omitempty"`
	UserId      bson.ObjectID              `validate:"required" bson:"user_id"`
	Name        string                     `validate:"required" bson:"name"`
	HitPoint    HitPointEntity             `validate:"required" bson:"hit_point"`
	CurrentExp  int32                      `validate:"required" bson:"current_exp"`
	AvatarImage string                     `validate:"required" bson:"avatar_image"`
	PocketMoney CharacterPocketMoneyEntity `validate:"required" bson:"pocket_money"`
	Proficiency CharacterProficiencyEntity `validate:"required" bson:"proficiencies"`
	Ability     AbilityEntity              `validate:"required" bson:"abilities"`
	ClassId     bson.ObjectID              `validate:"required" bson:"class_id"`
}

type CharacterEntityFilter struct {
	Id     *bson.ObjectID `bson:"_id,omitempty"`
	UserId *bson.ObjectID `bson:"user_id,omitempty"`
}
