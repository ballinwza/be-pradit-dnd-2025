package level_outbound_repository

import (
	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LevelRepository struct {
	collection *mongo.Collection
}

func NewLevelRepository() *LevelRepository {
	client := database.GetMongoClient()
	return &LevelRepository{
		collection: client.Database("pradit-dnd").Collection("levels"),
	}
}
