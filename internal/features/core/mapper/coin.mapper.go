package core_mapper

import (
	"fmt"

	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperCoinEntityToModel(entity core_outbound_entity.CoinEntity) *model.Coin {
	return &model.Coin{
		Name:      entity.Name,
		ShortType: mapperCoinShortTypeEntityToModel(entity.ShortType),
		Value:     *entity.Value,
	}
}
func MapperCoinModelToEntity(coinModel model.CoinReq) *core_outbound_entity.CoinEntity {
	return &core_outbound_entity.CoinEntity{
		Name:      mapperCoinShortTypeModelToEntityWithFullname(coinModel.ShortType),
		ShortType: mapperCoinShortTypeModelToEntity(coinModel.ShortType),
		Value:     &coinModel.Value,
	}
}

func mapperCoinShortTypeModelToEntityWithFullname(coinModel model.CoinShortType) string {
	switch coinModel {
	case model.CoinShortTypeCp:
		return "copper"
	case model.CoinShortTypeSp:
		return "silver"
	case model.CoinShortTypeEp:
		return "electrum"
	case model.CoinShortTypeGp:
		return "gold"
	case model.CoinShortTypePp:
		return "platinum"
	default:
		fmt.Printf("Error mapperCoinShortTypeModelToEntityWithFullname not match any type : %v", coinModel)
		return ""
	}
}

func mapperCoinShortTypeModelToEntity(coinModel model.CoinShortType) core_outbound_entity.CoinShortType {
	switch coinModel {
	case model.CoinShortTypeCp:
		return core_outbound_entity.CoinShortTypeCp
	case model.CoinShortTypeSp:
		return core_outbound_entity.CoinShortTypeSp
	case model.CoinShortTypeEp:
		return core_outbound_entity.CoinShortTypeEp
	case model.CoinShortTypeGp:
		return core_outbound_entity.CoinShortTypeGp
	case model.CoinShortTypePp:
		return core_outbound_entity.CoinShortTypePp
	default:
		return "unknow coinShortType"
	}
}
func mapperCoinShortTypeEntityToModel(entity core_outbound_entity.CoinShortType) model.CoinShortType {
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
