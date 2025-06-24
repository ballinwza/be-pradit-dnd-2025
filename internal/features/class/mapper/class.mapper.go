package class_mapper

import (
	"log"

	class_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/class/outbound/entity"
	core_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperClassEntityToModel(entity class_outbound_entity.ClassEntity) model.Class {
	afterMappingDice, err := core_mapper.MapperDiceRollTypeEntityToModel(entity.DiceHpIncrease)
	if err != nil {
		log.Println("MapperClassEntityToModel.MapperDiceRollTypeEntityToModel Error : ", err)
	}

	return model.Class{
		ID:             entity.Id.Hex(),
		Name:           entity.Name,
		DescriptionEn:  entity.DescriptionEn,
		DescriptionTh:  entity.DescriptionTh,
		DiceHpIncrease: afterMappingDice,
	}
}
