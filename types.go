package morm

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB represents the MongoDB client and database information.
type MongoDB struct {
	Client *mongo.Client
	DBName string
}

// Collect represents a MongoDB collection and model information.
type Collect struct {
	collection   *mongo.Collection
	modelType    reflect.Type
	modelElemPtr reflect.Value
}

// CollectQueryBuilder represents a query builder for MongoDB operations on a collection.
type CollectQueryBuilder struct {
	c          *Collect
	filter     interface{}
	skip       int64
	limit      int64
	projection bson.D
	sort       bson.D
	method     string
	popFields  []string
	value      interface{}
	pre        func(string, func())
}
