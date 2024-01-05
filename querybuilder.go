package morm

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// Skip sets the number of documents to skip
func (qb *CollectQueryBuilder) Skip(n int64) *CollectQueryBuilder {
	qb.skip = n
	return qb
}

// Limit sets the maximum number of documents to return
func (qb *CollectQueryBuilder) Limit(n int64) *CollectQueryBuilder {
	qb.limit = n
	return qb
}

// Projection sets the projection for the query
func (qb *CollectQueryBuilder) Projection(projection bson.D) *CollectQueryBuilder {
	qb.projection = projection
	return qb
}

// Sort sets the sort order for the query
func (qb *CollectQueryBuilder) Sort(sort bson.D) *CollectQueryBuilder {
	qb.sort = sort
	return qb
}

func (qb *CollectQueryBuilder) Populate(fields []string) *CollectQueryBuilder {
	qb.popFields = fields
	return qb
}

// Execute executes the query and returns the result
func (qb *CollectQueryBuilder) Exec() (interface{}, error) {
	result := reflect.New(qb.c.modelType.Elem()).Interface()

	if qb.method == "findone" {
		res, err := findone(qb, result)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	//Find
	res, err := find(qb, result)
	if err != nil {
		return nil, err
	}

	return res, nil
}
