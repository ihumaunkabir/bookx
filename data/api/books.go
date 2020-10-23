package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Book holds struct for book information
type Book struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Author    string             `json:"author,omitempty" bson:"author,omitempty"`
	Publisher string             `json:"psher,omitempty" bson:"psher,omitempty"`
	Year      string             `json:"year,omitempty" bson:"year,omitempty"`
	Category  string             `json:"cat,omitempty" bson:"cat,omitempty"`
	BookID    string             `json:"bookid,omitempty" bson:"bookid,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
