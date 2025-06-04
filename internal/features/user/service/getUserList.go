package user_service

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *UserService) GetUserList(ctx context.Context) ([]*model.User, error){
	var result []*model.User
	users, err := s.userRepository.FindAll(context.TODO())
	if err != nil {
		log.Fatal("Error fetching userList:", err)
	}

	for _, user := range users {
		result = append(result, MapperEntityToModel(&user))
	}

	return result, nil
}