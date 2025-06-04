package user_service

import (
	user_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperEntityToModel(entity *user_entity.UserEntity) *model.User {

	return &model.User{
		ID: entity.Id.Hex(),
		Email: entity.Email,
		DisplayName: entity.DisplayName,
		Password: entity.Password,
		UserImage: entity.UserImage,
	}
}