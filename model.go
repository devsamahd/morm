package morm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model represents a base MongoDB document model with common fields like ID, CreatedAt, and UpdatedAt.
// It is embedded in other structs to provide standardized fields for MongoDB documents.
type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
