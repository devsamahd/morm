package morm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create inserts a new document into the specified collection
func (qb *CollectQueryBuilder) Create(model interface{}, ctx ...context.Context) (primitive.ObjectID, error) {
	collection := qb.c.collection
	var backgroundContext = context.Background()
	if len(ctx) > 0 {
		backgroundContext = ctx[0]
	}

	res, err := collection.InsertOne(backgroundContext, model)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
