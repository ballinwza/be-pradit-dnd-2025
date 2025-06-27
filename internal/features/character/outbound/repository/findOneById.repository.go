package character_outbound_repository

import (
	"context"
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *CharacterRepository) FindOneById(ctx context.Context, objId bson.ObjectID) (*character_outbound_entity.CharacterEntity, error) {
	pipeline := mongo.Pipeline{
		bson.D{bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "_id", Value: objId},
		}}},
		// bson.D{
		// 	bson.E{Key: "$lookup", Value: bson.D{
		// 		bson.E{Key: "from", Value: "users"},
		// 		bson.E{Key: "localField", Value: "user_id"},
		// 		bson.E{Key: "foreignField", Value: "_id"},
		// 		bson.E{Key: "as", Value: "user_id"},
		// 	}},
		// },
		// bson.D{bson.E{Key: "$unwind", Value: "$user_id"}},
		// bson.D{
		// 	bson.E{Key: "$limit", Value: 1},
		// },
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("CharacterRepository.FindOneById Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Println("CharacterRepository.FindOneById Error : ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var result character_outbound_entity.CharacterEntity
		if err := cursor.Decode(&result); err != nil {
			log.Println("CharacterRepository.FindOneById decode error : ", err)
			return nil, err
		}
		// log.Println("Character : ", result)
		return &result, nil
	}

	return nil, err
}
