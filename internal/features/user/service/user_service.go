package user_service

import user_repository "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/repository"

type UserService struct {
	userRepository *user_repository.UserRepository
}

func NewUserService() *UserService {
	userRepository := user_repository.NewUserRepository()

	return &UserService{
		userRepository: userRepository,
	}
}