package morm

import (
	"testing"

	"github.com/devsamahd/morm"
)

func TestExex(t *testing.T) {
	// Connect to MongoDB for testing
	uri := "mongodb://localhost:27017"
	dbName := "test_db"
	_, err := morm.Connect(uri, dbName)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collectionName := "test_collection"
	model := &TestModel{}

	collect, err := morm.Collection(collectionName, model)
	if err != nil {
		t.Fatalf("Failed to create Collection: %v", err)
	}

	res, err := collect.Find().Skip(1).Limit(6).Populate([]string{""}).Exec()
	if err != nil {
		t.Fatalf("Exec failed to Execute Query: %v", err)
	}

	if res != 5 {

	}
}
