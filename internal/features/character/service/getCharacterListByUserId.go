package character_service

import (
	"context"
	"log"

	character_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *CharacterService) GetCharacterListByUserId(ctx context.Context, userId string) ([]*model.Character, error) {
	objId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Error converting Id to Hex at CharacterService.GetCharacterListByUserId : %v", err)
		return nil, err
	}
	characterEntity, err := s.characterRepository.FindAllByUserId(ctx, objId)
	if err != nil {
		log.Printf("Error CharacterService.GetCharacterListByUserId : %v", err)
		return nil, err
	}

	var characterModelList []*model.Character
	for _, value := range characterEntity {
		afterMapping := character_mapper.MapperCharacterEntityToModel(*value)
		characterModelList = append(characterModelList, &afterMapping)
	}
	return characterModelList, nil
}
