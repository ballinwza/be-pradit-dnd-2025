package character_repository

import (
	"context"
	"log"

	character_outbound_entity "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/outbound/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (r *CharacterRepository) WatchById(ctx context.Context, objId bson.ObjectID) (<-chan *character_outbound_entity.CharacterEntity, error) {
	characterEntityChan := make(chan *character_outbound_entity.CharacterEntity)

	go func() {
		defer close(characterEntityChan)
		log.Printf("CharacterRepository : Starting Change Stream for character ID : %s", objId.String())

		//Initial
		initialCharacter, err := r.FindOneById(ctx, objId)
		if err != nil {
			log.Println("CharacterRepository.WatchById Error : ", err)
		} else if initialCharacter != nil {
			log.Printf("CharacterRepository was initiated : %s", objId.String())
			characterEntityChan <- initialCharacter
		}

		pipeline := mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.D{
				{Key: "operationType", Value: bson.D{
					{Key: "$in", Value: bson.A{"insert", "update", "replace", "delete"}},
				}},
				{Key: "fullDocument._id", Value: objId},
			}}},
		}

		options := options.ChangeStream().SetFullDocument(options.UpdateLookup)
		changeStream, err := r.collection.Watch(ctx, pipeline, options)
		if err != nil {
			log.Println("CharacterRepository.WatchById Error : ", err)
			return
		}

		defer changeStream.Close(context.Background())

		// Start Watching
		log.Println("Starting to listen for Character changes...")
		for changeStream.Next(ctx) {
			var changeEvent struct {
				FullDocument character_outbound_entity.CharacterEntity `bson:"fullDocument"`
			}

			if err := changeStream.Decode(&changeEvent); err != nil {
				log.Println("CharacterRepository.WatchById decode Error : ", err)
				continue
			}

			log.Printf("Change detected for character: %s", changeEvent.FullDocument.Id.Hex())
			characterEntityChan <- &changeEvent.FullDocument
		}

		log.Printf("CharacterRepository.WatchById stop change streming !!")
	}()

	return characterEntityChan, nil
}
