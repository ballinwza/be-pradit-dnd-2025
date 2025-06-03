package user_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	user_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/outbound/user/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct{
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	client := database.GetMongoClient()

	return &UserRepository{
		collection: client.Database("pradit-dnd").Collection("users"),
	}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]user_entity.UserEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []user_entity.UserEntity
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindOneById(ctx context.Context, objId bson.ObjectID) (*user_entity.UserEntity, error) {
	var result user_entity.UserEntity
	err := r.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Not found document")
		} else {
			log.Fatalf("Find error: %v", err)
		}
		return nil, err
	}

	return &result, nil
}