package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AWSConfig struct {
	Region     string
	UserPoolID string
	ClientID   string
}

func LoadConfig() AWSConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	return AWSConfig{
		Region:     os.Getenv("AWS_REGION"),
		UserPoolID: os.Getenv("COGNITO_USER_POOL_ID"),
		ClientID:   os.Getenv("COGNITO_APP_CLIENT_ID"),
	}
}
