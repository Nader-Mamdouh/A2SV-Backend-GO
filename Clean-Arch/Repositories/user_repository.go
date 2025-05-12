package repositories

import (
	domain "JWT/Domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository defines the interface for user data access.
type UserRepository interface {
	CreateUser(user *domain.User) error
	FindUserByEmail(email string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (*domain.User, error)
}

// MongoUserRepository implements UserRepository using MongoDB.
type MongoUserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new MongoUserRepository instance.
func NewUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: collection}
}

// CreateUser creates a new user in MongoDB.
func (r *MongoUserRepository) CreateUser(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

// FindUserByEmail finds a user by email in MongoDB.
func (r *MongoUserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), map[string]string{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user in MongoDB.
func (r *MongoUserRepository) UpdateUser(user *domain.User) error {
	_, err := r.collection.UpdateOne(context.Background(), map[string]string{"id": user.ID}, user)
	return err
}

// GetAllUsers retrieves all users from MongoDB.
func (r *MongoUserRepository) GetAllUsers() ([]domain.User, error) {
	cursor, err := r.collection.Find(context.Background(), map[string]string{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []domain.User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a user by ID from MongoDB.
func (r *MongoUserRepository) GetUserByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), map[string]string{"id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
