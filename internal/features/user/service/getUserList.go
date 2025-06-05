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
		log.Println("UserService.GetUserList Error : ", err)
		return nil, err
	}

	for _, user := range users {
		afterMapping := user_mapper.MapperUserEntityToModel(user)
		result = append(result, &afterMapping)
	}

	return result, nil
}