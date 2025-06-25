package character_outbound_entity

import core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"

type AbilityEntity struct {
	Name      string                           `bson:"name"`
	Value     int32                            `bson:"value"`
	ShortType core_outbound_entity.AbilityType `bson:"short_type"`
}
