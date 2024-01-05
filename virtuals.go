package morm

import (
	"context"
	"reflect"
	"strings"

	"github.com/devsamahd/morm/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (qb *CollectQueryBuilder) Virtual(fields []string, value interface{}, filter interface{}) (interface{}, error) {
	modelType, err := utils.GetModelType(value)
	if err != nil {
		return nil, err
	}
	modelInstance := reflect.New(modelType).Interface()
	collectionName := strings.ToLower(modelType.Name()) + "s"

	var pipelineStages []bson.D

	pipelineStages = append(pipelineStages, bson.D{{Key: "$match", Value: filter}})

	for _, field := range fields {
		tags, err := utils.GetTags(modelType, field)
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
