package equipment_mapper

import (
	armor_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/mapper"
	equipment_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/outbound/entity"
	weapon_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperEquipmentEntityToModel(entity equipment_outbound_entity.EquipmentEntity) model.Equipment {
	afterMappingArmor := armor_mapper.MapperArmorEntityToModel(entity.Armor)
	afterMappingRightHanded := weapon_mapper.MapperWeaponEntityToModel(entity.RightHanded)
	afterMappingLeftHanded := weapon_mapper.MapperWeaponEntityToModel(entity.LeftHanded)
	return model.Equipment{
		ID:          entity.Id.Hex(),
		CharacterID: entity.CharacterId.Hex(),
		Armor:       &afterMappingArmor,
		RightHanded: &afterMappingRightHanded,
		LeftHanded:  &afterMappingLeftHanded,
	}
}

func MapEquipmentEntityToGqlModel(entity *equipment_outbound_entity.WatchEquipmentEntity) *model.WatchEquipment {
	id := entity.Id.Hex()
	cId := entity.CharacterId.Hex()
	aId := entity.ArmorId.Hex()
	rId := entity.RightHandedId.Hex()
	lId := entity.LeftHandedId.Hex()

	return &model.WatchEquipment{
		ID:          &id,
		CharacterID: &cId,
		Armor:       &aId,
		RightHanded: &rId,
		LeftHanded:  &lId,
	}
}
