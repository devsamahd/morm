package morm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Aggregate performs an aggregation query on the MongoDB collection using the specified pipeline.
// It takes a pipeline as an interface{}, representing the aggregation stages.
//
// Parameters:
//   - pipeline: The interface{} representing the aggregation pipeline stages.
//   - ctx: Optional context.Context for the aggregation operation. If not provided, the default context will be used.
//
// Returns:
//   - *mongo.Cursor: A cursor pointing to the result of the aggregation query.
//   - error: An error if any occurred during the aggregation operation.
func (c *Collect) Aggregate(pipeline interface{}, ctx ...context.Context) (*mongo.Cursor, error) {
	var backgroundContext = context.Background()
	if len(ctx) > 0 {
		backgroundContext = ctx[0]
	}
	return c.collection.Aggregate(backgroundContext, pipeline)
}
