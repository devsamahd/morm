package morm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Collect) Aggregate(pipeline interface{}, ctx ...context.Context) (*mongo.Cursor, error) {
	var backgroundContext = context.Background()
	if len(ctx) > 0 {
		backgroundContext = ctx[0]
	}
	return c.collection.Aggregate(backgroundContext, pipeline)
}