package character_outbound_entity

import core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"

type ProficiencyEntity struct {
	Name              core_outbound_entity.ProficiencyType `bson:"name"`
	Value             int32                                `bson:"value"`
	AbilityShortGroup core_outbound_entity.AbilityType     `bson:"ability_short_group"`
}
