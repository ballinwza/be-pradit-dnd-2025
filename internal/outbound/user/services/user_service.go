package user_service

import (
	"context"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	user_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/user/entities"
)

type UserService struct {}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context) (*model.User, error){
	mock := user_entity.UserEntity{
		Name: "test",
	}

	return &model.User{
		Email: mock.Name,
	}, nil
}