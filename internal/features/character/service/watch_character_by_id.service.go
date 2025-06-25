package character_service

import (
	"context"
	"log"

	character_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *CharacterService) WatchCharacterById(ctx context.Context, id string) (<-chan *model.Character, error) {
	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Println("CharacterService.WatchCharacterById Invalid ObjectId : ", err)
		return nil, err
	}

	repoChan, err := s.characterRepository.WatchById(ctx, objId)
	if err != nil {
		log.Println("CharacterService.WatchCharacterById Invalid ObjectId : ", err)
		return nil, err
	}

	characterModelChan := make(chan *model.Character)

	go func() {
		defer close(characterModelChan)

		for character := range repoChan {
			if character == nil {
				log.Println("CharacterService.WatchCharacterById: Received nil character from channel, skipping.")
				continue
			}

			afterMapping := character_mapper.MapperCharacterEntityToModel(*character)
			characterModelChan <- &afterMapping
		}
	}()

	return characterModelChan, nil
}
