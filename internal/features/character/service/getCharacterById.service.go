package character_service

import (
	"context"
	"log"

	character_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)



func (s *CharacterService) GetCharacterById(ctx context.Context, id string) (*model.Character, error){
	obj, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("CharacterService.GetCharacterById Invalid ObjectId : %v\n", err)
		return nil, err
	}

	character, err := s.characterRepository.FindOneById(context.Background(), obj)
	if err != nil {
		log.Printf("CharacterService.GetCharacterById Error : %v\n", err)
		return nil, err
	}

	afterMapping := character_mapper.MapperCharacterEntityToModel(*character)
	return &afterMapping, nil
}