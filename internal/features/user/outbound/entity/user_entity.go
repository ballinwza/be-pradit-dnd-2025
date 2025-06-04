package user_entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	Id bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	DisplayName string `bson:"display_name" json:"display_name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	UserImage string `bson:"user_image" json:"user_image"`
}