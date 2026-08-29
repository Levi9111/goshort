package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in our GoShort system.
type User struct {
	// primitive.ObjectID is Go's equivalent to mongoose.Types.ObjectId.
	// "omitempty" tells the MongoDB driver: "If this field is empty, don't include it when inserting".
	// This lets MongoDB auto-generate the _id for us on insert.

	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	// We use standard Go strings. 
	// Struct fields must be exported (capitalized) for BSON and JSON serializers to access them.
	// 1. bson: How it saves in MongoDB. 
	// 2. json: How it formats if we send this struct back to the frontend via an API response.
	Email string `bson:"email" json:"email"`

	// We use a hyphen "-" for json so the password is NEVER accidentally sent to the frontend.
	Password string `bson:"password" json:"-"`

	// Go's time.Time is equivalent to JS's Date object.
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}