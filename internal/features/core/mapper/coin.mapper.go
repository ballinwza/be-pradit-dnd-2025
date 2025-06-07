package core_mapper

import (
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCoinEntityToModel(entity core_outbound_entity.CoinEntity) model.Coin {
	return model.Coin{
		Name:      entity.Name,
		ShortType: mapCoinShortTypeEntityToModel(entity.ShortType),
		Value:     entity.Value,
	}
}

func mapCoinShortTypeEntityToModel(entity core_outbound_entity.CoinShortType) model.CoinShortType {
	switch entity {
	case core_outbound_entity.CoinShortTypeCp:
		return model.CoinShortTypeCp
	case core_outbound_entity.CoinShortTypeSp:
		return model.CoinShortTypeSp
	case core_outbound_entity.CoinShortTypeEp:
		return model.CoinShortTypeEp
	case core_outbound_entity.CoinShortTypeGp:
		return model.CoinShortTypeGp
	case core_outbound_entity.CoinShortTypePp:
		return model.CoinShortTypePp
	default:
		return "unknow coinShortType"
	}
}
