package repositories

import (
	"context"
	"errors"

	// Import our models package. Update the URL to match your go.mod!
	"github.com/levi9111/goshort/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	// 1. The Interface: The contract for what a User Repository must do.
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User,error)
}

// 2. The Struct: Holds our MongoDB collection pointer.
// It starts with a lowercase 'm', so it's private to this package.

type mongoUserRepository struct {
	collection *mongo.Collection
}

// 3. The Constructor: A normal function that returns our struct, but typed as the Interface.
func NewUserRepository(db *mongo.Database) UserRepository {
	return &mongoUserRepository{
		collection: db.Collection("users"),
	}
}


// 4. Receiver Functions: This is how Go does "methods" on a "class".
// The `(r *mongoUserRepository)` part means this function belongs to our struct.

// Create inserts a new user into the database.

func (r *mongoUserRepository) Create(ctx context.Context, user *models.User) error {
	// TS Equivalent: await collection.insertOne(user)
	_, err := r.collection.InsertOne(ctx, user) 
	return err
}

// FindByEmail searches for a user by their email address.

func (r *mongoUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	// bson.M is basically a map representing a JSON object.
	// TS Equivalent: { email: email }
	filter := bson.M{"email": email}

	// TS Equivalent: await collection.findOne(filter)
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		// If MongoDB says "no document found", we return a custom Go error
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Nil error, but nil user means "Not Found"
		}
		// If it's a real database crash, return the error
		return nil, err
	}

	return &user, nil
}