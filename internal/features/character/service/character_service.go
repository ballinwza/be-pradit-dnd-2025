package character_service

import character_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/repository"

type CharacterService struct {
	characterRepository *character_repository.CharacterRepository
}

func NewCharacterService() *CharacterService{
	characterRepository := character_repository.NewCharacterRepository()

	return &CharacterService{
		characterRepository: characterRepository,
	}
}