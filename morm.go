// Package morm provides a simple MongoDB wrapper for Go applications.
package morm

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBInstance represents a global instance of MongoDB connection.
var MongoDBInstance *MongoDB

// Connect establishes a connection to MongoDB and returns a MongoDB instance.
// It takes the MongoDB URI and the name of the database as parameters.
func Connect(uri string, dbName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Successfully connected to MongoDB")
	MongoDBInstance = &MongoDB{Client: client, DBName: dbName}
	return MongoDBInstance, nil
}

// Collection creates a new Collect instance for the specified collection and model.
// It takes the collection name and a model (pointer to a struct) as parameters.
// Returns a CollectQueryBuilder for building queries on the collection.
func Collection(collectionName string, model interface{}) (*CollectQueryBuilder, error) {
	collection := MongoDBInstance.Client.Database(MongoDBInstance.DBName).Collection(collectionName)
	modelType := reflect.TypeOf(model)

	if modelType.Kind() != reflect.Ptr || modelType.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("model must be a pointer to a struct")
	}

	modelElemPtr := reflect.New(modelType.Elem())

	c := &Collect{
		collection:   collection,
		modelType:    modelType,
		modelElemPtr: modelElemPtr,
	}

	return &CollectQueryBuilder{c: c}, nil
}
