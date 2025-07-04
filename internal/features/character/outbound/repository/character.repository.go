package character_outbound_repository

import (
	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CharacterRepository struct {
	collection *mongo.Collection
	validate   *validator.Validate
}

func NewCharacterRepository() *CharacterRepository {
	client := database.GetMongoClient()

	return &CharacterRepository{
		collection: client.Database("pradit-dnd").Collection("characters"),
		validate:   validator.New(),
	}
}
