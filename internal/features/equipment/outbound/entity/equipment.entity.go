package equipment_outbound_entity

import (
	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	weapon_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EquipmentEntity struct {
	Id            bson.ObjectID                       `bson:"_id,omitempty" json:"_id"`
	CharacterId   bson.ObjectID                       `bson:"character_id" json:"character_id"`
	ArmorId       bson.ObjectID                       `bson:"armor_id" json:"armor_id"`
	Armor         armor_outbound_entity.ArmorEntity   `bson:"armor" json:"armor"`
	RightHandedId bson.ObjectID                       `bson:"right_handed_id" json:"right_handed_id"`
	RightHanded   weapon_outbound_entity.WeaponEntity `bson:"right_handed" json:"right_handed"`
	LeftHandedId  bson.ObjectID                       `bson:"left_handed_id" json:"left_handed_id"`
	LeftHanded    weapon_outbound_entity.WeaponEntity `bson:"left_handed" json:"left_handed"`
}
