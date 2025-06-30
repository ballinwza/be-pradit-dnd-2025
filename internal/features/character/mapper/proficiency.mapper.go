package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterProficiencyEntityToModel(entity character_outbound_entity.CharacterProficiencyEntity) model.CharacterProficiency {
	return model.CharacterProficiency{
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
func MapperCharacterProficiencyModelToEntity(proficiencyModel model.CharacterProficiencyReq) *character_outbound_entity.CharacterProficiencyEntity {
	return &character_outbound_entity.CharacterProficiencyEntity{
		Athletics:      proficiencyModel.Athletics,
		Arobatics:      proficiencyModel.Arobatics,
		SleightOfHand:  proficiencyModel.SleightOfHand,
		Stealth:        proficiencyModel.Stealth,
		Arcana:         proficiencyModel.Arcana,
		History:        proficiencyModel.History,
		Investigation:  proficiencyModel.Investigation,
		Nature:         proficiencyModel.Nature,
		Religion:       proficiencyModel.Religion,
		AnimalHandling: proficiencyModel.AnimalHandling,
		Insight:        proficiencyModel.Insight,
		Medicine:       proficiencyModel.Medicine,
		Perception:     proficiencyModel.Perception,
		Survival:       proficiencyModel.Survival,
		Deception:      proficiencyModel.Deception,
		Intimidation:   proficiencyModel.Intimidation,
		Performance:    proficiencyModel.Performance,
		Persuasion:     proficiencyModel.Persuasion,
	}
}
