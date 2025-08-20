package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/roadguard/user_service/auth"
	"github.com/yourusername/roadguard/user_service/handlers"
)

func main() {
	cfg := auth.LoadConfig()

	cognitoSvc, err := auth.NewCognitoService(cfg)
	if err != nil {
		log.Fatal("Failed to init Cognito service:", err)
	}

	userHandler := handlers.NewUserHandler(cognitoSvc)

	r := gin.Default()
	r.POST("/signup", userHandler.Signup)
	r.POST("/login", userHandler.Login)

	log.Println("🚀 User Service running on :8080")
	r.Run(":8080")
}
