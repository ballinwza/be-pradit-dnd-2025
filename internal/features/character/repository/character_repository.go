package character_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	character_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (r *CharacterRepository) FindOneById(ctx context.Context, objId bson.ObjectID) (*character_entity.CharacterEntity, error) {
	pipeline := mongo.Pipeline{
		bson.D{bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "_id", Value: objId},
		}}},
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "users"},
				bson.E{Key: "localField", Value: "user_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "user_id"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$user_id"}},
		bson.D{
			bson.E{Key: "$limit", Value: 1},
		},
	}
	
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println("Error Character FindOneById : ", err)
	}
	defer cursor.Close(ctx)
	
	if cursor.Next(ctx) {
		var result character_entity.CharacterEntity
		if err := cursor.Decode(&result); err != nil {
			log.Println("Decode error:", err)
			return nil, fmt.Errorf("Decode error: %w", err)
		}
		return &result, nil
	}

	log.Println("No character found with id: ", objId.Hex(), pipeline)
	return nil, mongo.ErrNoDocuments
}