package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
}
