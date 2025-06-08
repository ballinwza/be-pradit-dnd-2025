package armor_outbound_repository

import (
	"context"
	"log"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/database"
	armor_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ArmorRepository struct {
	collection *mongo.Collection
}

func NewArmorRepository() *ArmorRepository {
	client := database.GetMongoClient()

	return &ArmorRepository{
		collection: client.Database("pradit-dnd").Collection("armors"),
	}
}

func (r *ArmorRepository) FindOneById(ctx context.Context, objId bson.ObjectID) (*armor_outbound_entity.ArmorEntity, error) {
	var result armor_outbound_entity.ArmorEntity
	err := r.collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("ArmorRepository.FindOneById Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Printf("ArmorRepository.FindOneById Error : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *ArmorRepository) FindAll(ctx context.Context, equipmentType armor_outbound_entity.ArmorEquipmentTypeArmorEntity) ([]*armor_outbound_entity.ArmorEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"armor_equipment_type": equipmentType})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("ArmorRepository.FindAll Error : Not founded document")
			return nil, mongo.ErrNoDocuments
		}

		log.Println("ArmorRepository.FindAll Error : ", err)
		return nil, err
	}

	var result []*armor_outbound_entity.ArmorEntity
	if err := cursor.All(ctx, &result); err != nil {
		log.Println("ArmorRepository.FindAll Error : ", err)
		return nil, err
	}

	return result, nil
}
