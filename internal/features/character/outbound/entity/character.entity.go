package character_outbound_entity

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CharacterEntity struct {
	Id          *bson.ObjectID                     `bson:"_id,omitempty"`
	Name        string                             `bson:"name"`
	HitPoint    HitPointEntity                     `bson:"hit_point"`
	CurrentExp  int32                              `bson:"current_exp"`
	AvatarImage string                             `bson:"avatar_image"`
	PocketMoney []*core_outbound_entity.CoinEntity `bson:"pocket_money"`
	Proficiency []*CharacterProficiencyEntity      `bson:"proficiencies"`
	Ability     []*AbilityEntity                   `bson:"abilities"`
	ClassId     bson.ObjectID                      `bson:"class_id"`
}
