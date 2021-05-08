package source

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database
var client *mongo.Client

func NewMongoClient() *mongo.Client {
	uri := Config.GetString("mongo")
	client, err := mongo.Connect(NewCtx(), options.Client().ApplyURI(uri).SetMaxPoolSize(20))
	if err != nil {
		logger.Fatal(err)
	}
	err = client.Ping(NewCtx(), readpref.Primary())
	if err != nil {
		logger.Fatal(err)
	} else {
		logger.Info("connect to database succeed")
	}
	return client
}

func getDatabase() *mongo.Database {
	db := GetMongoClient().Database("pancake")
	return db
}

func GetMongoClient() *mongo.Client {
	if client == nil {
		client = NewMongoClient()
	}
	return client
}
