package source

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func getMongoClient() *mongo.Client {
	uri := Config.GetString("mongo")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20))
	if err != nil {
		logger.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Fatal(err)
	} else {
		logger.Info("connect to database succeed")
	}
	return client
}

func getDatabase() *mongo.Database {
	c := getMongoClient()
	db := c.Database("pancake")
	return db
}
