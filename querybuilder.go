package morm

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Collect) Query() *CollectQueryBuilder {
	return &CollectQueryBuilder{c: c, filter: bson.M{}}
}

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
    collection := qb.c.collection

    result := reflect.New(qb.c.modelType.Elem()).Interface()

    if qb.method == "findone" {
        options := options.FindOne()
        if qb.projection != nil {
            options.SetProjection(qb.projection)
        }
        if qb.sort != nil {
            options.SetSort(qb.sort)
        }
        if qb.skip != 0 {
            options.SetSkip(qb.skip)
        }
        if qb.popFields != nil {
            result := reflect.New(qb.c.modelType.Elem()).Interface()
			resp, err := qb.Virtual(qb.popFields, result, qb.filter)
			if err != nil{
				return nil, err
			}
			return resp, nil
		}
	
        err := collection.FindOne(context.Background(), qb.filter, options).Decode(result)
		if err != nil {
			return nil, err
		}

        return result, nil
    }

	//Find
    options := options.Find()
    if qb.projection != nil {
        options.SetProjection(qb.projection)
    }
    if qb.sort != nil {
        options.SetSort(qb.sort)
    }
    if qb.skip != 0 {
        options.SetSkip(qb.skip)
    }
    if qb.limit != 0 {
        options.SetLimit(qb.limit)
    }
    
    cursor, err := collection.Find(context.Background(), qb.filter, options)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var results []interface{}

    for cursor.Next(context.Background()) {
        if qb.popFields != nil {
            result := reflect.New(qb.c.modelType.Elem()).Interface()
            err := cursor.Decode(result)
            if err != nil {
                return nil, err
            }

            id := reflect.ValueOf(result).Elem().FieldByName("ID")
            popRes, err := qb.Virtual(qb.popFields, result, bson.M{"_id": id.Interface().(primitive.ObjectID)})
            if err != nil{
                return nil, err
            }
            results = append(results, popRes)
        }else{
            result := reflect.New(qb.c.modelType.Elem()).Interface()
            err := cursor.Decode(result)
            if err != nil {
                return nil, err
            }
            results = append(results, result)
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }
	
	
    return results, nil
}
