package user_outbound_entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	Id    bson.ObjectID `bson:"_id,omitempty"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
	Image string        `bson:"image"`
}
