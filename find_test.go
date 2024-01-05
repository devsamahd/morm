package morm

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockMongoDB is a mock for MongoDB to be used in testing
type MockMongoDB struct {
	collection *mongo.Collection
}

// Mock the Connect function for testing
func MockConnect(uri, dbName string) (*MockMongoDB, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MockMongoDB{
		collection: client.Database(dbName).Collection("test_collection"),
	}, nil
}


// TestFind tests the Find method
func TestFind(t *testing.T) {
	// Mock the MongoDB connection
	mockDB, err := MockConnect("mongodb://localhost:27017", "test_db")
	if err != nil {
		t.Fatalf("Failed to mock MongoDB connection: %v", err)
	}

	// Mock the collection
	model := &TestModel{}
	qb, err := Collection("test_collection", model)
	if err != nil {
		t.Fatalf("Failed to mock collection: %v", err)
	}

	// Add some data to the collection for testing
	mockDB.collection.InsertOne(context.Background(), bson.M{"Field1": "value", "Field2": 42})

	// Test the Find method
	result, err := qb.Find(bson.M{"Field1": "value"}).Exec()
	if err != nil {
		t.Fatalf("Find method returned an error: %v", err)
	}
	println(result)

	// Perform assertions on the result if needed
	// ...

	// Clean up (delete the test data)
	mockDB.collection.DeleteMany(context.Background(), bson.M{"Field1": "value"})
}

// TestFindOne tests the FindOne method
func TestFindOne(t *testing.T) {
	mockDB, err := MockConnect("mongodb://localhost:27017", "test_db")
	if err != nil {
		t.Fatalf("Failed to mock MongoDB connection: %v", err)
	}

	// Mock the collection
	model := &TestModel{}
	qb, err := Collection("test_collection", model)
	if err != nil {
		t.Fatalf("Failed to mock collection: %v", err)
	}

	// Test the FindOne method
	result, err := qb.FindOne(bson.M{"Field1": "value"}).Exec()
	if err != nil {
		t.Fatalf("FindOne method returned an error: %v", err)
	}

	println(result)
	// Perform assertions on the result if needed
	// ...

	// Clean up (delete the test data)
	mockDB.collection.DeleteMany(context.Background(), bson.M{"Field1": "value"})
}

// TestFindOneAndUpdate tests the FindOneAndUpdate method
func TestFindOneAndUpdate(t *testing.T) {
	mockDB, err := MockConnect("mongodb://localhost:27017", "test_db")
	if err != nil {
		t.Fatalf("Failed to mock MongoDB connection: %v", err)
	}

	// Mock the collection
	model := &TestModel{}
	qb, err := Collection("test_collection", model)
	if err != nil {
		t.Fatalf("Failed to mock collection: %v", err)
	}

	// Test the FindOneAndUpdate method
	update := bson.M{"$set": bson.M{"Field2": 100}}
	result, err := qb.FindOneAndUpdate(bson.M{"Field1": "value"}, update)
	if err != nil {
		t.Fatalf("FindOneAndUpdate method returned an error: %v", err)
	}
	println(result)

	// Perform assertions on the result if needed
	// ...

	// Clean up (delete the test data)
	mockDB.collection.DeleteMany(context.Background(), bson.M{"Field1": "value"})
}

// TestFindOneAndRemove tests the FindOneAndRemove method
func TestFindOneAndRemove(t *testing.T) {
	mockDB, err := MockConnect("mongodb://localhost:27017", "test_db")
	if err != nil {
		t.Fatalf("Failed to mock MongoDB connection: %v", err)
	}

	// Mock the collection
	model := &TestModel{}
	qb, err := Collection("test_collection", model)
	if err != nil {
		t.Fatalf("Failed to mock collection: %v", err)
	}

	// Test the FindOneAndRemove method
	result, err := qb.FindOneAndRemove(bson.M{"Field1": "value"})
	if err != nil {
		t.Fatalf("FindOneAndRemove method returned an error: %v", err)
	}

	// Perform assertions on the result if needed
	// ...

	// Clean up (delete the test data)
	// Note: You may need to check if result is nil or not before deleting
	if result != nil {
		mockDB.collection.DeleteMany(context.Background(), bson.M{"Field1": "value"})
	}
}
