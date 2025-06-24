package class_service

import (
	"context"
	"log"

	class_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/class/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *ClassService) GetClassList(ctx context.Context) ([]*model.Class, error) {
	var result []*model.Class

	classModel, err := s.classRepository.FindAll(ctx)
	if err != nil {
		log.Println("ClassService.GetClassList Error : ", err)
	}

	for _, class := range classModel {
		afterMapping := class_mapper.MapperClassEntityToModel(*class)
		result = append(result, &afterMapping)
	}

	return result, nil
}
