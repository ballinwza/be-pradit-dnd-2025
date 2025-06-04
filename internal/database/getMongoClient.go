package database

import (
	"sync"

	"github.com/ballinwza/be-pradit-dnd-2025/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	clientInstance *mongo.Client
	once 	sync.Once
)

func GetMongoClient() (*mongo.Client) {
	once.Do(func() {
		config.LoadingEnv()
		mongoURI := config.GetMongoURI()

		clientInstance = connectMongo(mongoURI)
	})

	return clientInstance
}
