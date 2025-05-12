package controllers

import (
	"net/http"

	domain "JWT/Domain"
	usecases "JWT/Usecases"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related HTTP requests.
type AuthController struct {
	userUsecase *usecases.UserUsecase
}

// NewAuthController creates a new AuthController instance.
func NewAuthController(userUsecase *usecases.UserUsecase) *AuthController {
	return &AuthController{userUsecase: userUsecase}
}

// Register handles user registration.
func (ac *AuthController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.UserType == "" {
		user.UserType = "user"
	}

	err := ac.userUsecase.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user login.
func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.userUsecase.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"role":       user.UserType,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
	})
}

// GetAllUsers retrieves all users.
func (ac *AuthController) GetAllUsers(c *gin.Context) {
	users, err := ac.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser retrieves a user by ID.
func (ac *AuthController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := ac.userUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"role":       user.UserType,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})
}
