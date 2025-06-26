package level_outbound_repository

import (
	"context"
	"log"

	level_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/level/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *LevelRepository) FindAll(ctx context.Context) ([]*level_outbound_entity.LevelEntity, error) {
	filter := bson.M{}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("LevelRepository.FindAll Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Println("LevelRepository.FindAll Error : ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var levelList []*level_outbound_entity.LevelEntity
	if err := cursor.All(ctx, &levelList); err != nil {
		log.Println("LevelRepository.FindAll Error : ", err)
		return nil, err
	}

	return levelList, nil
}
