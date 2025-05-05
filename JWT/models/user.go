package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID  `bson:"_id"`
	FirstName *string             `json:"first_name" validate:"required" min:"2" max:"100"`
	LastName  *string             `json:"last_name" validate:"required" min:"2" max:"100"`
	Email     *string             `json:"email" validate:"required,email"`
	Password  *string             `json:"password" validate:"required,min=6,max=100"`
	User_type *string             `json:"role" validate:"required, eq=admin|eq=user"`
	Phone     *string             `json:"phone" validate:"required,min=10,max=15"`
	CreatedAt *primitive.DateTime `bson:"created_at"`
	UpdatedAt *primitive.DateTime `bson:"updated_at"`
	User_id   *string             `json:"user_id" validate:"required"`
	Token	  *string             `json:"token"`
	RefreshToken *string          `json:"refresh_token"`
}