package weapon_mapper

import (
	weapon_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperMasteryEntityToModel(entity weapon_outbound_entity.WeaponMasteryEntity) model.WeaponMastery {
	return model.WeaponMastery{
		ID:            entity.Id.Hex(),
		Name:          entity.Name,
		DescriptionEn: entity.DescriptionEn,
		DescriptionTh: entity.DescriptionTh,
		ImageURL:      entity.ImageUrl,
	}
}
