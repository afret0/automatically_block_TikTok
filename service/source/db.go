package source

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var db *mongo.Database

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

func GetDatabase() *mongo.Database {
	if db == nil {
		db = GetMongoClient().Database("pancake")
	}
	return db
}

func GetMongoClient() *mongo.Client {
	if client == nil {
		client = NewMongoClient()
	}
	return client
}
