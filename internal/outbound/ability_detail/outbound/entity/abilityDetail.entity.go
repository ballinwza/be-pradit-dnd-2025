package ability_detail_outbound_entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AbilityDetailShortTypeEntity string

const (
	AbilityDetailShortTypeEntityStr AbilityDetailShortTypeEntity = "str"
	AbilityDetailShortTypeEntityDex AbilityDetailShortTypeEntity = "dex"
	AbilityDetailShortTypeEntityCon AbilityDetailShortTypeEntity = "con"
	AbilityDetailShortTypeEntityInt AbilityDetailShortTypeEntity = "int"
	AbilityDetailShortTypeEntityWis AbilityDetailShortTypeEntity = "wis"
	AbilityDetailShortTypeEntityCha AbilityDetailShortTypeEntity = "cha"
)

type AbilityDetailEntity struct {
	Id bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name" json:"name"`
	Short AbilityDetailShortTypeEntity `bson:"short" json:"short"`
	DescriptionEn string `bson:"description_en" json:"description_en"`
	DescriptionTh string `bson:"description_th" json:"description_th"`
	Proficiencies []ProficiencyDetailEntity `bson:"proficiencies" json:"proficiencies"`
}

