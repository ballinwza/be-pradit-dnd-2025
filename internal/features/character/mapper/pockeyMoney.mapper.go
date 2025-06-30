package character_mapper

import (
	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCharacterPocketMoneyEntityToModel(entity character_outbound_entity.CharacterPocketMoneyEntity) *model.CharacterPocketMoney {
	return &model.CharacterPocketMoney{
		Cp: &entity.Cp,
		Sp: &entity.Sp,
		Ep: &entity.Ep,
		Gp: &entity.Gp,
		Pp: &entity.Pp,
	}
}
func MapperCharacterPocketMoneyModelToEntity(coinModel model.CharacterPocketMoneyReq) *character_outbound_entity.CharacterPocketMoneyEntity {
	return &character_outbound_entity.CharacterPocketMoneyEntity{
		Cp: *coinModel.Cp,
		Sp: *coinModel.Sp,
		Ep: *coinModel.Ep,
		Gp: *coinModel.Gp,
		Pp: *coinModel.Pp,
	}
}
