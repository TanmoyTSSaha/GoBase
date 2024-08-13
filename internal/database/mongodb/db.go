package mongodb

import (
	"context"
	"log"

	"github.com/TanmoyTSSaha/GoBase/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func MongoConnect(url, dbName string, config *configs.Configuration) (*MongoDBClient, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	logsDBExists := true
	if !config.Internals.IsLogDependenciesCreated {
		logsDBExists, err = checkDBExists(dbName, client)
		if err != nil {
			return nil, err
		}
	}

	if !logsDBExists {
		createDatabase(dbName, client)
		updatedConfig := config
		updatedConfig.Internals.IsLogDependenciesCreated = true
		configs.UpdateCongif(updatedConfig)
	}

	database := client.Database(dbName)

	return &MongoDBClient{Client: client, Database: database}, nil
}

func checkDBExists(dbName string, client *mongo.Client) (bool, error) {
	dbExists := false

	databases, err := client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		return dbExists, err
	}

	for _, db := range databases {
		if dbName == db {
			dbExists = true
			break
		}
	}

	return dbExists, nil
}

func createDatabase(dbName string, client *mongo.Client) error {
	db := client.Database(dbName)
	collection := db.Collection("init_collection")
	_, err := collection.InsertOne(context.Background(), bson.M{"status": "initialized"})
	if err != nil {
		return err
	}

	log.Printf("DATABASE '%s' CREATED SUCCESSFULLY", dbName)

	return nil
}
