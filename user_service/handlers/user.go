package handlers

import (
	"log"
	"net/http"

	"github.com/Satyaprakash2507/RoadGuard/user_service/auth"
	"github.com/Satyaprakash2507/RoadGuard/user_service/models"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related routes
type UserHandler struct {
	Cognito *auth.CognitoService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(cognito *auth.CognitoService) *UserHandler {
	return &UserHandler{Cognito: cognito}
}

// Signup creates a new user in Cognito
func (h *UserHandler) Signup(c *gin.Context) {
	var req models.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("❌ Signup validation failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Cognito.Signup(req.Email, req.Password); err != nil {
		log.Println("❌ Cognito signup failed:", err)
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "✅ Signup successful! Please confirm your email."})
}

// Login authenticates user and returns JWT tokens
func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("❌ Login validation failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	access, id, refresh, err := h.Cognito.Login(req.Email, req.Password)
	if err != nil {
		log.Println("❌ Cognito login failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"id_token":      id,
		"refresh_token": refresh,
	})
}
