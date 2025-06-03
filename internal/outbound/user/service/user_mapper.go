package user_service

import (
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	user_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/user/entity"
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