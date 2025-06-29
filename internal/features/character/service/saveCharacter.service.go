package character_service

import (
	"context"
	"log"

	character_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *CharacterService) SaveCharacter(ctx context.Context, characterModel *model.SaveCharacterReq) (bool, error) {
	afterMapping, err := character_mapper.MapperCharacterModelToEntity(*characterModel)
	if err != nil {
		log.Printf("Error CharacterService.SaveCharacter can't coverting Id : %v", err)
		return false, err
	}

	if characterModel.ID != nil {
		result, err := s.characterRepository.UpdateOneById(ctx, *afterMapping.Id, *afterMapping)
		if err != nil {
			log.Printf("Error CharacterService.SaveCharacter can't fetch characterRepository : %v", err)
			return false, err
		}

		return result, nil

	}

	result, err := s.characterRepository.InsertOne(ctx, *afterMapping)
	if err != nil {
		log.Printf("Error CharacterService.SaveCharacter can't fetch characterRepository : %v", err)
		return false, err
	}

	return result, nil
}
