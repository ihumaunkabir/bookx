package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Basic holds basic info for db record
type Basic struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
	Deleted   bool          `bson:"deleted"`
}

// PreparePut prepares Basic before database put operation
func (b *Basic) PreparePut() {
	b.UpdatedAt = time.Now().UTC()

	if b.ID == "" {
		b.ID = bson.NewObjectId()
		b.CreatedAt = b.UpdatedAt
	}
}

// PrepareDelete updates delete info
func (b *Basic) PrepareDelete() {
	b.Deleted = true
}
