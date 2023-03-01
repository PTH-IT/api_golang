package utils

import (
	"fmt"

	"github.com/PTH-IT/api_golang/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func NewSecretsManager() *secretsmanager.SecretsManager {
	awsEndpoint := fmt.Sprintf("%s:%s", config.Getconfig().Aws.Host, config.Getconfig().Aws.Port)
	awsRegion := config.Getconfig().Aws.Region
	id := config.Getconfig().Aws.Id
	secret := config.Getconfig().Aws.Secret
	token := config.Getconfig().Aws.Token
	svc := secretsmanager.New(session.New(&aws.Config{
		Region:           aws.String(awsRegion),
		Endpoint:         aws.String(awsEndpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(id, secret, token),
	}))
	return svc
}
func GetscretManager(secretName string) string {
	svc := NewSecretsManager()
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		fmt.Printf("failed to get secret manager %q, %v", secretName, err)
	}
	return *result.SecretString
}
