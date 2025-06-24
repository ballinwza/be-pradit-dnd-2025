package class_service

import class_outbound_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/class/outbound/repository"

type ClassService struct {
	classRepository *class_outbound_repository.ClassRepository
}

func NewClassService() *ClassService {
	classRepository := class_outbound_repository.NewClassRepository()

	return &ClassService{
		classRepository: classRepository,
	}
}
