package morm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToUpdateStruct(data interface{}) (primitive.M, error) {
	// Marshal the struct to BSON
	bsonData, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal the BSON to primitive.M
	var result primitive.M
	err = bson.Unmarshal(bsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateOne updates a single document in the specified collection based on the filter and update parameters\
func (qb *CollectQueryBuilder) UpdateOne(filter interface{}, update interface{}) error {
	collection := qb.c.collection

	convertedUpdate, _ := ToUpdateStruct(update)
	// Set updatedAt field to the current time
	updateWithUpdatedAt := bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	}

	// Merge the provided update with the update for updatedAt
	for key, value := range convertedUpdate {
		updateWithUpdatedAt["$set"].(bson.M)[key] = value
	}

	// Perform the update
	_, err := collection.UpdateOne(context.Background(), filter, updateWithUpdatedAt)
	return err
}

// Update updates multiple documents in the specified collection based on the filter and update parameters
func (qb *CollectQueryBuilder) Update(filter interface{}, update interface{}, ctx ...context.Context) error {
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

	_, err := collection.UpdateMany(backgroundContext, filter, updateWithUpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
