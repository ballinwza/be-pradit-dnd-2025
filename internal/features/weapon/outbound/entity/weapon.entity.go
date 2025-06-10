package weapon_outbound_entity

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type WeaponEntity struct {
	Id            bson.ObjectID                           `bson:"_id,omitempty" json:"_id"`
	Name          string                                  `bson:"name" json:"name"`
	DamagesType   core_outbound_entity.DamagedType        `bson:"damaged_type" json:"damaged_type"`
	DiceRollType  core_outbound_entity.DiceRollTypeEntity `bson:"dice_roll_type" json:"dice_roll_type"`
	DiceQuantity  int32                                   `bson:"dice_quantity" json:"dice_quantity"`
	Price         core_outbound_entity.CoinEntity         `bson:"price" json:"price"`
	DescriptionEn string                                  `bson:"description_en" json:"description_en"`
	DescriptionTh string                                  `bson:"description_th" json:"description_th"`
	Weight        core_outbound_entity.WeightEntity       `bson:"weight" json:"weight"`
	MasteryId     WeaponMasteryEntity                     `bson:"mastery_id" json:"mastery_id"`
	PropertyId    []WeaponPropertyEntity                  `bson:"property_id" json:"property_id"`
	ImageUrl      string                                  `bson:"image_url" json:"image_url"`
}
