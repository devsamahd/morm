package morm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToUpdateStruct converts a struct to a primitive.M for MongoDB update operations.
// It marshals the input struct to BSON and then unmarshals it to a primitive.M.
//
// Parameters:
//   - data: The struct to be converted.
//
// Returns:
//   - primitive.M: The converted BSON data.
//   - error: An error if the conversion fails.
//
// Example:
//   type UpdateData struct {
//     Field1 string `bson:"field1"`
//     Field2 int    `bson:"field2"`
//   }
//
//   data := UpdateData{"value1", 42}
//   converted, err := ToUpdateStruct(data)
//   if err != nil {
//     // Handle error
//   }
//
// This function is useful for converting a struct to the format required for MongoDB update operations.
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

// UpdateOne updates a single document in the specified collection based on the filter and update parameters.
//
// Parameters:
//   - filter: The filter criteria to identify the document to update.
//   - update: The update data to be applied to the document.
//
// Returns:
//   - error: An error if the update operation fails.
//
// Example:
//   filter := bson.D{{"name", "John"}}
//   update := bson.D{{"$set", bson.D{{"age", 30}}}}
//   err := qb.UpdateOne(filter, update)
//   if err != nil {
//     // Handle error
//   }
//
// This method is useful for updating a single document in a MongoDB collection.
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

// Update updates multiple documents in the specified collection based on the filter and update parameters.
//
// Parameters:
//   - filter: The filter criteria to identify the documents to update.
//   - update: The update data to be applied to the documents.
//   - ctx: Optional context for the MongoDB operation.
//
// Returns:
//   - error: An error if the update operation fails.
//
// Example:
//   filter := bson.D{{"status", "pending"}}
//   update := bson.D{{"$set", bson.D{{"status", "processed"}}}}
//   err := qb.Update(filter, update)
//   if err != nil {
//     // Handle error
//   }
//
// This method is useful for updating multiple documents in a MongoDB collection.
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
