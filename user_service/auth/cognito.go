package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// CognitoService wraps AWS Cognito client and config
type CognitoService struct {
	Client   *cognitoidentityprovider.Client
	ClientID string
}

// NewCognitoService initializes a new Cognito client with AWS configuration
func NewCognitoService(cfg AWSConfig) (*CognitoService, error) {
	// Load AWS config with region
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, fmt.Errorf("❌ failed to load AWS config: %w", err)
	}

	client := cognitoidentityprovider.NewFromConfig(awsCfg)

	return &CognitoService{
		Client:   client,
		ClientID: cfg.ClientID,
	}, nil
}

// Signup registers a new user in Cognito User Pool
func (c *CognitoService) Signup(email, password string) error {
	// Always use context with timeout for external calls
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.Client.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: &c.ClientID,
		Username: &email,
		Password: &password,
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: &email,
			},
		},
	})

	if err != nil {
		return fmt.Errorf("❌ signup failed for %s: %w", email, err)
	}
	return nil
}

// Login authenticates the user and returns JWT tokens
func (c *CognitoService) Login(email, password string) (string, string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.Client.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
		ClientId: &c.ClientID,
	})
	if err != nil {
		return "", "", "", fmt.Errorf("❌ login failed for %s: %w", email, err)
	}

	if resp.AuthenticationResult == nil {
		return "", "", "", fmt.Errorf("⚠️ no authentication result received for %s", email)
	}

	// Return AccessToken, IdToken, RefreshToken
	return aws.ToString(resp.AuthenticationResult.AccessToken),
		aws.ToString(resp.AuthenticationResult.IdToken),
		aws.ToString(resp.AuthenticationResult.RefreshToken),
		nil
}
