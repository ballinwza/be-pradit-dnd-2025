package character_outbound_repository

import (
	"context"
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	core_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/core/outbound/entity"
)

func checkType(firstLoopValue []*core_outbound_entity.CoinEntity, second []*core_outbound_entity.CoinEntity) {
	for _, value := range firstLoopValue {
		for index, pocketItem := range second {
			if pocketItem.ShortType == value.ShortType {
				second[index] = value
			}
		}
	}
}

func (r CharacterRepository) InsertOne(ctx context.Context, characterEntity character_outbound_entity.CharacterEntity) (bool, error) {
	defaultVal := 0.00
	var test *float64 = &defaultVal

	defaultCoin := []*core_outbound_entity.CoinEntity{
		&core_outbound_entity.CoinEntity{
			Name:      "copper",
			Value:     test,
			ShortType: core_outbound_entity.CoinShortTypeCp,
		},
		&core_outbound_entity.CoinEntity{
			Name:      "silver",
			Value:     test,
			ShortType: core_outbound_entity.CoinShortTypeSp,
		},
		&core_outbound_entity.CoinEntity{
			Name:      "electrum",
			Value:     test,
			ShortType: core_outbound_entity.CoinShortTypeEp,
		},
		&core_outbound_entity.CoinEntity{
			Name:      "gold",
			Value:     test,
			ShortType: core_outbound_entity.CoinShortTypeGp,
		},
		&core_outbound_entity.CoinEntity{
			Name:      "platinum",
			Value:     test,
			ShortType: core_outbound_entity.CoinShortTypePp,
		},
	}

	checkType(characterEntity.PocketMoney, defaultCoin)

	factory := character_outbound_entity.CharacterEntity{
		Name:        characterEntity.Name,
		HitPoint:    characterEntity.HitPoint,
		CurrentExp:  characterEntity.CurrentExp,
		AvatarImage: characterEntity.AvatarImage,
		PocketMoney: defaultCoin,
		ClassId:     characterEntity.ClassId,
	}

	result, err := r.collection.InsertOne(ctx, factory)
	if err != nil {
		log.Println("Error CharacterRepository.InsertOne : ", err)
		return false, err
	}

	return result.Acknowledged, nil
}
