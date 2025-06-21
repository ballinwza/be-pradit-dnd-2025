package character_mapper

import (
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperProficiencyEntityToModel(entity character_entity.ProficiencyEntity) model.Proficiency {

	return model.Proficiency{
		Athletics:      entity.Athletics,
		Arobatics:      entity.Arobatics,
		SleightOfHand:  entity.SleightOfHand,
		Stealth:        entity.Stealth,
		Arcana:         entity.Arcana,
		History:        entity.History,
		Investigation:  entity.Investigation,
		Nature:         entity.Nature,
		Religion:       entity.Religion,
		AnimalHandling: entity.AnimalHandling,
		Insight:        entity.Insight,
		Medicine:       entity.Medicine,
		Perception:     entity.Perception,
		Survival:       entity.Survival,
		Deception:      entity.Deception,
		Intimidation:   entity.Intimidation,
		Performance:    entity.Performance,
		Persuasion:     entity.Persuasion,
	}
}
