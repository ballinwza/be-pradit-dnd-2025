package class_outbound_repository

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	class_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/class/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClassRepository struct {
	collection *mongo.Collection
}

func NewClassRepository() *ClassRepository {
	client := database.GetMongoClient()

	return &ClassRepository{
		collection: client.Database("pradit-dnd").Collection("classes"),
	}
}

func (r *ClassRepository) FindAll(ctx context.Context) ([]*class_outbound_entity.ClassEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("ClassRepository.FindAll Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Println("ClassRepository.FindAll Error : ", err)
		return nil, err
	}

	var result []*class_outbound_entity.ClassEntity
	if err := cursor.All(ctx, &result); err != nil {
		log.Println("ClassRepository.FindAll Error : ", err)
		return nil, err
	}

	return result, nil
}
