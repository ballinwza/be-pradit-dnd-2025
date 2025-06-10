package weapon_mapper

import (
	weapon_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperPropertyEntityToModel(entity weapon_outbound_entity.WeaponPropertyEntity) model.WeaponProperty {
	return model.WeaponProperty{
		ID:            entity.Id.Hex(),
		Name:          entity.Name,
		DescriptionEn: entity.DescriptionEn,
		DescriptionTh: entity.DescriptionTh,
		ImageURL:      entity.ImageUrl,
	}
}
