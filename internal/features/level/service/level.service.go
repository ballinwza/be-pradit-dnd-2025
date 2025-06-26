package level_service

import level_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/level/outbound/repository"

type LevelService struct {
	levelRepository *level_outbound_repository.LevelRepository
}

func NewLevelService() *LevelService {
	levelRepository := level_outbound_repository.NewLevelRepository()
	return &LevelService{
		levelRepository: levelRepository,
	}
}
