package morm

import (
	"context"
	"testing"
)

// TestConnect tests the Connect function
func TestConnect(t *testing.T) {
	// Provide your MongoDB URI and database name for testing
	uri := "mongodb://localhost:27017"
	dbName := "test_db"

	_, err := Connect(uri, dbName)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check if MongoDBInstance is not nil
	if MongoDBInstance == nil {
		t.Fatal("MongoDBInstance should not be nil after successful connection")
	}

	// Check if the MongoDBInstance has the correct DBName
	if MongoDBInstance.DBName != dbName {
		t.Fatalf("Expected DBName %s, got %s", dbName, MongoDBInstance.DBName)
	}

	// Check if Ping is successful
	err = MongoDBInstance.Client.Ping(context.Background(), nil)
	if err != nil {
		t.Fatalf("Ping to MongoDB failed: %v", err)
	}
}

// TestCollection tests the Collection function
func TestCollection(t *testing.T) {
	// Connect to MongoDB for testing
	uri := "mongodb://localhost:27017"
	dbName := "test_db"
	_, err := Connect(uri, dbName)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collectionName := "test_collection"
	model := &TestModel{}

	collect, err := Collection(collectionName, model)
	if err != nil {
		t.Fatalf("Failed to create Collection: %v", err)
	}

	// Check if the collection is not nil
	if collect.Find() == nil {
		t.Fatal("Collection should not be nil")
	}
}
