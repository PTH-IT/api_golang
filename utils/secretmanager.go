package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetscretManager(secretName string) string {
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"
	svc := secretsmanager.New(session.New(&aws.Config{
		Region:           aws.String(awsRegion),
		Endpoint:         aws.String(awsEndpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials("hau", "hau", ""),
	}))
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
