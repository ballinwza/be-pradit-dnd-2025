package user_mapper

import (
	user_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
)

func MapperUserEntityToModel(entity user_outbound_entity.UserEntity) model.User {

	return model.User{
		ID:    entity.Id.Hex(),
		Email: entity.Email,
		Name:  entity.Name,
		Image: entity.Image,
	}
}
