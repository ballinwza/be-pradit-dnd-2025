package level_outbound_entity

import "go.mongodb.org/mongo-driver/v2/bson"

type LevelEntity struct {
	Id                    bson.ObjectID `bson:"_id,omitempty"`
	Level                 int32         `bson:"level"`
	ExpToLevelUp          int32         `bson:"exp_to_level_up"`
	ProficiencyBonusPoint int32         `bson:"proficiency_bonus_point"`
}
