package class_outbound_entity

import (
	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ClassType string

const (
	ClassTypeBabarian ClassType = "babarian"
	ClassTypeBard     ClassType = "bard"
	ClassTypeCleric   ClassType = "cleric"
	ClassTypeDruid    ClassType = "druid"
	ClassTypeFigther  ClassType = "figther"
	ClassTypeMonk     ClassType = "monk"
	ClassTypePaladin  ClassType = "paladin"
	ClassTypeRanger   ClassType = "ranger"
	ClassTypeRogue    ClassType = "rogue"
	ClassTypeSorcerer ClassType = "sorcerer"
	ClassTypeWarlock  ClassType = "warlock"
	ClassTypeWizard   ClassType = "wizard"
)

type ClassEntity struct {
	Id                     bson.ObjectID                           `bson:"_id,omitempty"`
	Name                   string                                  `bson:"name"`
	DescriptionEn          string                                  `bson:"description_en"`
	DescriptionTh          string                                  `bson:"description_th"`
	DiceHpIncrease         core_outbound_entity.DiceRollTypeEntity `bson:"dice_hp_increase"`
	FitProficiencyIds      []bson.ObjectID                         `bson:"fit_proficiency_ids"`
	FitArmorTypes          []armor_outbound_entity.ArmorTypeEntity `bson:"fit_armor_types"`
	FitWeaponProficiencies []string                                `bson:"fit_weapon_proficiencies"`
	StarterEquipments      []ClassStartEquipmentEntity             `bson:"starter_equipments"`
	// FeatureIds           []bson.ObjectID                         `bson:"feature_ids"`
}

type ClassStartEquipmentEntity struct {
	chooseType  string                            `bson:"choose_type"`
	pocketMoney []core_outbound_entity.CoinEntity `bson:"pocket_money"`
	items       []ClaseStartItemEntity            `bson:"items"`
}

type ClaseStartItemEntity struct {
	itemId   bson.ObjectID `bson:"item_id"`
	quantity int32         `bson:"quantity"`
}
