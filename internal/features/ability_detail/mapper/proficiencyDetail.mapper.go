package ability_detail_mapper

import (
	ability_detail_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperProficiencyDetailEntityToModel(entity ability_detail_outbound_entity.ProficiencyDetailEntity) model.ProficiencyDetail {
	return model.ProficiencyDetail{
		Name:          entity.Name,
		DescriptionEn: entity.DescriptionEn,
		DescriptionTh: entity.DescriptionTh,
	}
}
