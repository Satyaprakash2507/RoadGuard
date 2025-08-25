package main

import (
	"log"

	"github.com/Satyaprakash2507/RoadGuard/user_service/auth"
	"github.com/Satyaprakash2507/RoadGuard/user_service/handlers"
	"github.com/gin-gonic/gin"
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
