package character_outbound_repository

import (
	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CharacterRepository struct {
	collection *mongo.Collection
}

func NewCharacterRepository() *CharacterRepository {
	client := database.GetMongoClient()

	return &CharacterRepository{
		collection: client.Database("pradit-dnd").Collection("characters"),
	}
}
