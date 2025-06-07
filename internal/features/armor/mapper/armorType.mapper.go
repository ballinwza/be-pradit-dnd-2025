package armor_mapper

import (
	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperArmorTypeEntityToModel(entity armor_outbound_entity.ArmorTypeEntity) model.ArmorType {
	switch entity {
	case armor_outbound_entity.ArmorTypeEntityLight:
		return model.ArmorTypeLight
	case armor_outbound_entity.ArmorTypeEntityMedium:
		return model.ArmorTypeMedium
	case armor_outbound_entity.ArmorTypeEntityHeavy:
		return model.ArmorTypeHeavy
	default:
		return "unknow ArmorType"
	}
}
