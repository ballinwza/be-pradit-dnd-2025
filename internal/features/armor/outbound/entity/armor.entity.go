package armor_outbound_entity

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ArmorTypeEntity string

const (
	ArmorTypeEntityLight  ArmorTypeEntity = "light"
	ArmorTypeEntityMedium ArmorTypeEntity = "medium"
	ArmorTypeEntityHeavy  ArmorTypeEntity = "heavy"
)

type ArmorEntity struct {
	Id                    bson.ObjectID                            `bson:"_id,omitempty" json:"_id"`
	ArmorType             ArmorTypeEntity                          `bson:"armor_type" json:"armor_type"`
	Name                  string                                   `bson:"name" json:"name"`
	ArmorClass            int32                                    `bson:"armor_class" json:"armor_class"`
	MaximumPlusArmorClass *int32                                   `bson:"maximum_plus_ac,omitempty" json:"maximum_plus_ac"`
	StealthAdvantage      core_outbound_entity.AdvantageTypeEntity `bson:"stealth_advantage" json:"stealth_advantage"`
	Price                 core_outbound_entity.CoinEntity          `bson:"price" json:"price"`
	StrRequirement        *int32                                   `bson:"str_requirement,omitempty" json:"str_requirement"`
	DescriptionEn         string                                   `bson:"description_en" json:"description_en"`
	Weight                core_outbound_entity.WeightEntity        `bson:"weight" json:"weight"`
}
