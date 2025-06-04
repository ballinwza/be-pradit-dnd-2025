package character_entity

import (
	user_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CharacterEntity struct {
	Id bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name" json:"name"`
	AvatarImage string `bson:"avatar_image" json:"avatar_image"`
	Speed int32 `bson:"speed" json:"speed"`
	InitiativePlusPoint int32 `bson:"initiative_plus_point" json:"initiative_plus_point"`
	HitDice int32 `bson:"hit_dice" json:"hit_dice"`

	UserId user_entity.UserEntity `bson:"user_id" json:"user_id"`
}