package morm

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Client *mongo.Client
	DBName string
}

type Collect struct {
	collection   *mongo.Collection
	modelType    reflect.Type
	modelElemPtr reflect.Value
}

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
	pre		   func(string, func())
}