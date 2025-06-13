package equipment_outbound_repository

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	equipment_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EquipmentRepository struct {
	collection *mongo.Collection
}

func NewEquipmentRepository() *EquipmentRepository {
	client := database.GetMongoClient()

	return &EquipmentRepository{
		collection: client.Database("pradit-dnd").Collection("equipments"),
	}
}

// type EqpFilter struct {
// 	ChracterId bson.ObjectID `bson:"character_id"`
// }

func (r *EquipmentRepository) FindByCharacterId(ctx context.Context, objId bson.ObjectID) (*equipment_outbound_entity.EquipmentEntity, error) {
	// filter := EqpFilter{
	// 	ChracterId: bson.ObjectIDFromHex("6839b9ad34d38e31825f2ba2"),
	// }

	pipeline := mongo.Pipeline{
		bson.D{bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "character_id", Value: objId},
		}}},
		// Armor
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "armors"},
				bson.E{Key: "localField", Value: "armor_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "armor"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$armor"}},
		// Right Handed
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapons"},
				bson.E{Key: "localField", Value: "right_handed_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "right_handed"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$right_handed"}},
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_mastery"},
				bson.E{Key: "localField", Value: "right_handed.mastery_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "right_handed.mastery_id"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$right_handed.mastery_id"}},
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_property"},
				bson.E{Key: "localField", Value: "right_handed.property_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "right_handed.property_id"},
			}},
		},
		//Left Handed
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapons"},
				bson.E{Key: "localField", Value: "left_handed_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "left_handed"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$left_handed"}},
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_mastery"},
				bson.E{Key: "localField", Value: "left_handed.mastery_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "left_handed.mastery_id"},
			}},
		},
		bson.D{bson.E{Key: "$unwind", Value: "$left_handed.mastery_id"}},
		bson.D{
			bson.E{Key: "$lookup", Value: bson.D{
				bson.E{Key: "from", Value: "weapon_property"},
				bson.E{Key: "localField", Value: "left_handed.property_id"},
				bson.E{Key: "foreignField", Value: "_id"},
				bson.E{Key: "as", Value: "left_handed.property_id"},
			}},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("EquipmentRepository.FindByCharacterId Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Println("EquipmentRepository.FindByCharacterId Error : ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var result equipment_outbound_entity.EquipmentEntity
		if err := cursor.Decode(&result); err != nil {
			log.Println("EquipmentRepository.FindByCharacterId decode error : ", err)
			return nil, err
		}
		return &result, nil
	}

	return nil, err
}
