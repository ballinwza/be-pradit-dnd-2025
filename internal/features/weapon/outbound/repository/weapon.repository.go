package weapon_outbound_repository

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	weapon_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type WeaponRepository struct {
	collection *mongo.Collection
}

func NewWeaponRepository() *WeaponRepository {
	client := database.GetMongoClient()

	return &WeaponRepository{
		collection: client.Database("pradit-dnd").Collection("weapons"),
	}
}

func (r *WeaponRepository) FindAll(ctx context.Context) ([]*weapon_outbound_entity.WeaponEntity, error) {
	pipeline := mongo.Pipeline{
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_mastery"},
				bson.E{Key: "localField", Value: "mastery_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "mastery_id"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$mastery_id"}},
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_property"},
				bson.E{Key: "localField", Value: "property_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "property_id"},
			}},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("WeaponRepository.FindAll Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Println("WeaponRepository.FindAll Error : ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var result []*weapon_outbound_entity.WeaponEntity
	if err := cursor.All(ctx, &result); err != nil {
		log.Println("WeaponRepository.FindAll Error : ", err)
		return nil, err
	}

	return result, nil
}

func (r *WeaponRepository) FindById(ctx context.Context, id bson.ObjectID) (*weapon_outbound_entity.WeaponEntity, error) {
	pipeline := mongo.Pipeline{
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_mastery"},
				bson.E{Key: "localField", Value: "mastery_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "mastery_id"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$mastery_id"}},
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_property"},
				bson.E{Key: "localField", Value: "property_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "property_id"},
			}},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("WeaponRepository.FindAll Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Println("WeaponRepository.FindAll Error : ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var result *weapon_outbound_entity.WeaponEntity
		if err := cursor.Decode(&result); err != nil {
			log.Println("EquipmentRepository.FindByCharacterId decode error : ", err)
			return nil, err
		}
		return result, nil
	}

	return nil, err
}
