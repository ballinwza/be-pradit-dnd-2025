package user_service

import (
	"context"
	"log"

	user_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func (s *UserService) GetUserList(ctx context.Context) ([]*model.User, error){
	var result []*model.User
	users, err := s.userRepository.FindAll(context.TODO())
	if err != nil {
		log.Fatal("Error fetching userList:", err)
	}

	for _, user := range users {
		afterMapping := user_mapper.MapperUserEntityToModel(user)
		result = append(result, &afterMapping)
	}

	return result, nil
}