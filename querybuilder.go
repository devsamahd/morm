package morm

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// Skip sets the number of documents to skip in the MongoDB query.
// It specifies the number of documents to skip before starting to return documents.
//
// Parameters:
//   - n: The number of documents to skip.
//
// Example:
//
//	qb.Skip(10).Limit(5).Execute()
//
// This method is useful for implementing pagination in query results.
func (qb *CollectQueryBuilder) Skip(n int64) *CollectQueryBuilder {
	qb.skip = n
	return qb
}

// Limit sets the maximum number of documents to return in the MongoDB query.
// It specifies the maximum number of documents to return in the result set.
//
// Parameters:
//   - n: The maximum number of documents to return.
//
// Example:
//
//	qb.Limit(20).Sort(bson.D{{"timestamp", -1}}).Execute()
//
// This method is commonly used to control the size of the result set and manage query performance.
func (qb *CollectQueryBuilder) Limit(n int64) *CollectQueryBuilder {
	qb.limit = n
	return qb
}

// Projection sets the projection fields for the MongoDB query.
// It specifies the fields to include or exclude from the query result.
//
// Parameters:
//   - projection: A BSON document representing the projection.
//
// Example:
//
//	qb.Projection(bson.D{{"name", 1}, {"age", 1}}).Skip(5).Execute()
//
// This method is useful for selecting specific fields in the query result.
func (qb *CollectQueryBuilder) Projection(projection bson.D) *CollectQueryBuilder {
	qb.projection = projection
	return qb
}

// Sort sets the sort order for the MongoDB query.
// It specifies the order in which the query result should be sorted.
//
// Parameters:
//   - sort: A BSON document representing the sort order.
//
// Example:
//
//	qb.Sort(bson.D{{"timestamp", -1}}).Limit(10).Execute()
//
// This method is commonly used for ordering query results based on one or more fields.
func (qb *CollectQueryBuilder) Sort(sort bson.D) *CollectQueryBuilder {
	qb.sort = sort
	return qb
}

// Populate sets the fields to populate in the MongoDB query result.
// It takes a slice of strings representing the field names to populate.
//
// Example:
//
//	qb.Populate([]string{"Author", "Comments"})
//
// This method is commonly used to specify fields that are references to other collections
// and need to be populated with their actual values in the query result.
//
// Note: Ensure that the provided field names are valid and exist in the MongoDB documents.
func (qb *CollectQueryBuilder) Populate(fields []string) *CollectQueryBuilder {
	qb.popFields = fields
	return qb
}

// Exec executes the MongoDB query and returns the result.
// It performs the specified MongoDB operation (e.g., find, findOne) based on the query configuration.
//
// Returns:
//   - result: A pointer to the result model where the query result will be decoded.
//   - error: An error if the query execution fails.
//
// Example:
//
//	var user User
//	err := qb.Filter(bson.D{{"name", "John"}}).Exec(&user)
//	if err != nil {
//	  // Handle error
//	}
//
// This method is used to execute the configured MongoDB query and retrieve the result.
// The type of 'result' should be a pointer to the model struct representing the collection documents.
// The actual operation performed depends on the query method set using the Method() method.
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
