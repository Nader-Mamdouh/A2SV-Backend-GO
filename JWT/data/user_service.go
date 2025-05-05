package data

import (
	"context"
	"errors"
	"log"
	"time"

	"JWT/models"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection
var jwtKey = os.Getenv("JWT_SECRET")
func ConnectDB() {
	MONGO_DB:= os.Getenv("MONGO_URL")
	clientOption := options.Client().ApplyURI(MONGO_DB)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	db := os.Getenv("DB_NAME")               // Define the database name
	collectionname := "users" // Define the collection name
	UserCollection = client.Database(db).Collection(collectionname)
}

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{collection: collection}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Check if user already exists
	var existingUser models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	password := string(hashedPassword)
	user.Password = &password

	// Set timestamps
	now := primitive.NewDateTimeFromTime(time.Now())
	user.CreatedAt = &now
	user.UpdatedAt = &now

	// Generate user ID
	user.ID = primitive.NewObjectID()
	userID := user.ID.Hex()
	user.User_id = &userID

	// Insert user
	_, err = s.collection.InsertOne(context.Background(), user)
	return err
}

func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (s *UserService) GenerateTokens(user *models.User) (string, string, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.User_id,
		"role":    user.User_type,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(jwtKey)) 
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.User_id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(jwtKey)) 
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
