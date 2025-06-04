package user_service

import (
	"context"
	"log"

	user_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/mapper"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)



func (s *UserService) GetUserById(ctx context.Context, id string) (*model.User, error){
	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("Invalid ObjectId: %v", err)
	}
	
	user, err := s.userRepository.FindOneById(context.Background(), objId)
	if err != nil {
		log.Fatal("Error fetching user:", err)
	}

	afterMapping := user_mapper.MapperUserEntityToModel(*user)
	return &afterMapping, nil
}

