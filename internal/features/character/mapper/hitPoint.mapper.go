package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperHitPointEntityToModel(entity character_outbound_entity.HitPointEntity) model.HitPoint {

	return model.HitPoint{
		MaxHp:          entity.MaxHp,
		CurrentHp:      entity.CurrentHp,
		TemporaryHp:    entity.TemporaryHp,
		MaxTemporaryHp: entity.MaxTemporaryHp,
	}
}
func MapperHitPointModelToEntity(hpModel model.HitPointReq) character_outbound_entity.HitPointEntity {
	return character_outbound_entity.HitPointEntity{
		MaxHp:          *hpModel.MaxHp,
		CurrentHp:      *hpModel.CurrentHp,
		TemporaryHp:    *hpModel.TemporaryHp,
		MaxTemporaryHp: *hpModel.MaxTemporaryHp,
	}
}
