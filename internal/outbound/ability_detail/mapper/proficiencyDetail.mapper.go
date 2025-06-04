package ability_detail_mapper

import (
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	ability_detail_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/ability_detail/outbound/entity"
)

func MapperProficiencyDetailEntityToModel(entity ability_detail_outbound_entity.ProficiencyDetailEntity) model.ProficiencyDetail{
	return model.ProficiencyDetail{
		Name: entity.Name,
		DescriptionEn: entity.DescriptionEn,
		DescriptionTh: entity.DescriptionTh,
	}
}
