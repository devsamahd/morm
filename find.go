package morm

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/devsamahd/morm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (qb *CollectQueryBuilder) Find(filter ...interface{}) *CollectQueryBuilder {
    if len(filter) > 0 {
        qb.filter = filter[0]
	}
	qb.method = "find"
	return qb
}

func (qb *CollectQueryBuilder) FindOne(filter interface{}) *CollectQueryBuilder {
	qb.filter = filter
	qb.method = "findone"
	return qb
}

func (qb *CollectQueryBuilder) FindOneAndUpdate(filter interface{}, update interface{}, ctx ...context.Context) (interface{}, error) {
	collection := qb.c.collection
	var backgroundContext = context.Background()
	if len(ctx) > 0 {
		backgroundContext = ctx[0]
	}

	// Set updatedAt field to the current time
	updateWithUpdatedAt := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	}

	// Merge the provided update with the update for updatedAt
	updateWithUpdatedAt["$set"].(bson.M)["$set"] = update

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	result := collection.FindOneAndUpdate(backgroundContext, filter, updateWithUpdatedAt, options)
	if result.Err() != nil {
		return nil, result.Err()
	}

	// Decode the result into the original model
	modelType, err := utils.GetModelType(qb.c.modelElemPtr.Interface())
	if err != nil {
		return nil, err
	}

	resultValue := reflect.New(modelType).Interface()
	err = result.Decode(resultValue)
	if err != nil {
		return nil, err
	}

	return resultValue, nil
}

// FindOneAndRemove finds a single document in the specified collection based on the filter and removes it
func (qb *CollectQueryBuilder) FindOneAndRemove(filter interface{}, ctx ...context.Context) (interface{}, error) {
	collection := qb.c.collection
	var backgroundContext = context.Background()
	if len(ctx) > 0 {
		backgroundContext = ctx[0]
	}

	options := options.FindOneAndDelete().SetProjection(bson.D{{Key: "_id", Value: 0}})

	result := collection.FindOneAndDelete(backgroundContext, filter, options)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, result.Err()
	}

	// Decode the result into the original model
	modelType, err := utils.GetModelType(qb.c.modelElemPtr.Interface())
	if err != nil {
		return nil, err
	}

	resultValue := reflect.New(modelType).Interface()
	err = result.Decode(resultValue)
	if err != nil {
		return nil, err
	}

	return resultValue, nil
}
