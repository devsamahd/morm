package morm

import "context"

// Delete deletes a document from the specified collection based on the filter
func (qb *CollectQueryBuilder) Delete(filter interface{}, ctx ...context.Context) error {
	collection := qb.c.collection
	var backgroundContext = context.Background()
	if len(ctx) > 0 {
		backgroundContext = ctx[0]
	}

	_, err := collection.DeleteOne(backgroundContext, filter)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMany deletes multiple documents from the specified collection based on the filter
func (qb *CollectQueryBuilder) DeleteMany(filter interface{}, ctx ...context.Context) (int64, error) {
	collection := qb.c.collection
	var backgroundContext = context.Background()
	if len(ctx) > 0 {
		backgroundContext = ctx[0]
	}

	result, err := collection.DeleteMany(backgroundContext, filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
