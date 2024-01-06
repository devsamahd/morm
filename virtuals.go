package morm

import (
	"context"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// Virtual performs a virtual lookup on the specified fields and updates the provided value accordingly.
// It uses the MongoDB aggregation framework to perform the lookup and populate the specified fields.
//
// The virtual lookup involves creating a $lookup stage for each field, linking to the related collection,
// and updating the value based on the specified local and foreign fields. If the field has "justOne" set to true,
// it uses $arrayElemAt to ensure a single value is returned.
//
// Parameters:
//   - fields: A slice of strings specifying the fields to perform virtual lookup.
//   - value: The interface{} value to be updated with the virtual lookup results.
//   - filter: The filter to match documents for the virtual lookup.
//
// Returns:
//   - interface{}: The updated value after the virtual lookup.
//   - error: An error if any occurred during the virtual lookup process.
func (qb *CollectQueryBuilder) virtual(fields []string, value interface{}, filter interface{}) (interface{}, error) {
	modelType, err := getModelType(value)
	if err != nil {
		return nil, err
	}
	modelInstance := reflect.New(modelType).Interface()
	collectionName := strings.ToLower(modelType.Name()) + "s"

	var pipelineStages []bson.D

	pipelineStages = append(pipelineStages, bson.D{{Key: "$match", Value: filter}})

	for _, field := range fields {
		tags, err := getTags(modelType, field)
		if err != nil {
			return nil, err
		}

		localField := tags["localField"]
		foreignField := tags["foreignField"]
		justOne := tags["justOne"]
		count := tags["count"]

		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: strings.ToLower(field) + "s"},
				{Key: "localField", Value: localField},
				{Key: "foreignField", Value: foreignField},
				{Key: "as", Value: strings.ToLower(field)},
			}},
		}

		pipelineStages = append(pipelineStages, lookupStage)

		if count != "true" && justOne == "true" {
			// Use $ifNull to conditionally set the virtual field based on whether it exists or not
			pipelineStages = append(pipelineStages, bson.D{
				{Key: "$addFields", Value: bson.M{
					strings.ToLower(field): bson.M{"$arrayElemAt": bson.A{"$" + strings.ToLower(field), 0}},
				}},
			})
		}
	}

	collection, err := Collection(collectionName, modelInstance)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.c.Aggregate(pipelineStages)
	if err != nil {
		return nil, err
	}

	// Check if the virtual document exists
	if cursor.Next(context.Background()) {
		// Decode the virtual document
		if err := cursor.Decode(value); err != nil {
			return nil, err
		}
	}

	return value, nil
}
