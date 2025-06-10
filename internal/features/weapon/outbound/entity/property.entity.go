package weapon_outbound_entity

import "go.mongodb.org/mongo-driver/v2/bson"

type WeaponPropertyEntity struct {
	Id            bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name          string        `bson:"name" json:"name"`
	DescriptionEn string        `bson:"description_en" json:"description_en"`
	DescriptionTh string        `bson:"description_th" json:"description_th"`
	ImageUrl      string        `bson:"image_url" json:"image_url"`
}
