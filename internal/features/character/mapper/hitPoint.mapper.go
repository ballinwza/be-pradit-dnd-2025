package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperHitPointEntityToModel(entity character_outbound_entity.HitPointEntity) model.HitPoint {

	return model.HitPoint{
		MaxHp:          entity.MaxHp,
		CurrrentHp:     entity.CurrentHp,
		TemporaryHp:    entity.TemporaryHp,
		MaxTemporaryHp: entity.MaxTemporaryHp,
	}
}
