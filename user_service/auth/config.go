package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AWSConfig struct {
	Region     string
	UserPoolID string
	ClientID   string
}

// LoadConfig loads AWS Cognito config from .env or system env vars
// and validates required fields.
func LoadConfig() AWSConfig {
	// 1. Try to load local .env (for dev only)
	_ = godotenv.Load() // ignore error, we’ll check vars anyway

	// 2. Fetch from environment
	cfg := AWSConfig{
		Region:     os.Getenv("AWS_REGION"),
		UserPoolID: os.Getenv("COGNITO_USER_POOL_ID"),
		ClientID:   os.Getenv("COGNITO_APP_CLIENT_ID"),
	}

	// 3. Validate
	missing := []string{}
	if cfg.Region == "" {
		missing = append(missing, "AWS_REGION")
	}
	if cfg.UserPoolID == "" {
		missing = append(missing, "COGNITO_USER_POOL_ID")
	}
	if cfg.ClientID == "" {
		missing = append(missing, "COGNITO_APP_CLIENT_ID")
	}

	if len(missing) > 0 {
		log.Fatal(fmt.Errorf("❌ Missing required environment variables: %v", missing))
	}

	return cfg
}
