package auth

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoService struct {
	Client   *cognitoidentityprovider.Client
	ClientID string
}

func NewCognitoService(cfg AWSConfig) (*CognitoService, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.Region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := cognitoidentityprovider.NewFromConfig(awsCfg)

	return &CognitoService{
		Client:   client,
		ClientID: cfg.ClientID,
	}, nil
}

func (c *CognitoService) Signup(email, password string) error {
	_, err := c.Client.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: &c.ClientID,
		Username: &email,
		Password: &password,
		UserAttributes: []cognitoidentityprovider.AttributeType{
			{
				Name:  awsString("email"),
				Value: &email,
			},
		},
	})
	return err
}

func (c *CognitoService) Login(email, password string) (string, string, error) {
	resp, err := c.Client.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
		ClientId: &c.ClientID,
	})
	if err != nil {
		return "", "", err
	}
	return *resp.AuthenticationResult.AccessToken, *resp.AuthenticationResult.IdToken, nil
}

func awsString(s string) *string {
	return &s
}
