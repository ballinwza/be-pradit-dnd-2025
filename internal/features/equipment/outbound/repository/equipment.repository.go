package equipment_outbound_repository

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	equipment_mapper "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/mapper"
	equipment_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/outbound/entity"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (r *EquipmentRepository) FindByCharacterId(ctx context.Context, objId bson.ObjectID) (*equipment_outbound_entity.EquipmentEntity, error) {

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

func (r *EquipmentRepository) WatchByCharacterId(ctx context.Context, objId bson.ObjectID) (<-chan *model.Equipment, error) {
	equipmentChan := make(chan *model.Equipment)

	go func() {
		defer close(equipmentChan)
		log.Printf("EquipmentRepository : Starting Change Stream for equipment ID : %s", objId.String())

		//Initial
		initialEquipment, err := r.FindByCharacterId(ctx, objId)
		if err != nil {
			log.Println("EquipmentRepository.WatchByCharacterId Error : ", err)
		} else if initialEquipment != nil {
			log.Printf("EquipmentRepository : InitialEquipment : %s", objId.String())
			resultMapping := equipment_mapper.MapperEquipmentEntityToModel(*initialEquipment)
			equipmentChan <- &resultMapping
		}

		pipeline := mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.D{
				{Key: "operationType", Value: bson.D{
					{Key: "$in", Value: bson.A{"insert", "update", "replace"}},
				}},
				{Key: "fullDocument.character_id", Value: objId},
			}}},
		}

		changeStream, err := r.collection.Watch(ctx, pipeline, options.ChangeStream().SetFullDocument(options.UpdateLookup))
		if err != nil {
			log.Println("EquipmentRepository.WatchByCharacterId Error : ", err)
			return
		}

		defer changeStream.Close(context.Background())

		log.Println("Starting to listen for equipment changes...")
		for changeStream.Next(ctx) {
			var changeEvent struct {
				FullDocument equipment_outbound_entity.WatchEquipmentEntity `bson:"fullDocument"`
			}

			if err := changeStream.Decode(&changeEvent); err != nil {
				log.Println("EquipmentRepository.WatchByCharacterId decode Error : ", err)
				continue
			}

			completeEquipment, err := r.FindByCharacterId(ctx, changeEvent.FullDocument.CharacterId)
			if err != nil {
				log.Println("EquipmentRepository.WatchByCharacterId Error : ", err)
				continue
			}

			if completeEquipment == nil {
				log.Println("EquipmentRepository.WatchByCharacterId return Nil equipment")
				continue
			}

			log.Printf("Change detected for character: %s", changeEvent.FullDocument.CharacterId.Hex())
			resultMapping := equipment_mapper.MapperEquipmentEntityToModel(*completeEquipment)
			equipmentChan <- &resultMapping

		}

		log.Printf("EquipmentRepository.WatchByCharacterId stop streming !!")
	}()

	return equipmentChan, nil

}
