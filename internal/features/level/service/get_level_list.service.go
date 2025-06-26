package level_service

import (
	"context"
	"log"

	level_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/level/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *LevelService) GetLevelList(ctx context.Context) ([]*model.Level, error) {
	var levelListModel []*model.Level
	levelListEntity, err := s.levelRepository.FindAll(ctx)
	if err != nil {
		log.Printf("LevelService.GetLevelList Error : %v\n", err)
		return nil, err
	}

	for _, value := range levelListEntity {
		afterMapping := level_mapper.MapperLevelEntityToModel(*value)
		levelListModel = append(levelListModel, &afterMapping)
	}

	return levelListModel, nil
}
