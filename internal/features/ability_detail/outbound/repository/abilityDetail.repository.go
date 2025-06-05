package ability_detail_outbound_repository

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	ability_detail_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/outbound/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AbilityDetailRepository struct{
	collection *mongo.Collection
}

func NewAbilityDetailRepository() *AbilityDetailRepository {
	client:= database.GetMongoClient()

	return &AbilityDetailRepository{
		collection: client.Database("pradit-dnd").Collection("ability_detail"),
	}
}

func (r *AbilityDetailRepository) FindOneByShort(ctx context.Context, short ability_detail_entity.AbilityDetailShortTypeEntity)(*ability_detail_entity.AbilityDetailEntity, error){
	var result ability_detail_entity.AbilityDetailEntity
	err := r.collection.FindOne(ctx, bson.M{"short": short}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("AbilityDetailRepository.FindOneByShort Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		} 

		log.Printf("AbilityDetailRepository.FindOneByShort Error : %v\n", err)
		return nil, err
	}

	return &result, nil
}