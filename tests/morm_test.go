package morm

import (
	"context"
	"testing"

	"github.com/devsamahd/morm"
)

// TestConnect tests the Connect function
func TestConnect(t *testing.T) {
	// Provide your MongoDB URI and database name for testing
	uri := "mongodb://localhost:27017"
	dbName := "test_db"

	_, err := morm.Connect(uri, dbName)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check if MongoDBInstance is not nil
	if morm.MongoDBInstance == nil {
		t.Fatal("MongoDBInstance should not be nil after successful connection")
	}

	// Check if the MongoDBInstance has the correct DBName
	if morm.MongoDBInstance.DBName != dbName {
		t.Fatalf("Expected DBName %s, got %s", dbName, morm.MongoDBInstance.DBName)
	}

	// Check if Ping is successful
	err = morm.MongoDBInstance.Client.Ping(context.Background(), nil)
	if err != nil {
		t.Fatalf("Ping to MongoDB failed: %v", err)
	}
}

// TestCollection tests the Collection function
func TestCollection(t *testing.T) {
	// Connect to MongoDB for testing
	uri := "mongodb://localhost:27017"
	dbName := "test_db"
	_, err := morm.Connect(uri, dbName)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Test with a valid model
	type TestModel struct {
		ID   string `bson:"_id,omitempty"`
		Name string `bson:"name"`
	}

	collectionName := "test_collection"
	model := &TestModel{}

	collect, err := morm.Collection(collectionName, model)
	if err != nil {
		t.Fatalf("Failed to create Collection: %v", err)
	}

	// Check if the collection is not nil
	if collect.Query() == nil {
		t.Fatal("Collection should not be nil")
	}
}
