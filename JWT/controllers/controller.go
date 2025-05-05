package controllers

import (
	"net/http"

	"JWT/data"
	"JWT/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthController struct {
	userService *data.UserService
}

func NewAuthController(userService *data.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default role if not provided
	if user.User_type == nil {
		defaultRole := "user"
		user.User_type = &defaultRole
	}

	err := ac.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.userService.AuthenticateUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, refreshToken, err := ac.userService.GenerateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Update user with tokens
	user.Token = &accessToken
	user.RefreshToken = &refreshToken

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":         user.User_id,
			"email":      user.Email,
			"role":       user.User_type,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
	})
}
func GetUser (c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	err = data.UserCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.User_id,
		"email":      user.Email,
		"role":       user.User_type,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})
}
func GetAllUsers(c *gin.Context) {
	cursor, err := data.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(context.Background())

	var users []gin.H
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			continue
		}

		users = append(users, gin.H{
			"id":         user.User_id,
			"email":      user.Email,
			"role":       user.User_type,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}