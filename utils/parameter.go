package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetParameter(name string) (string, error) {
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:           aws.String(awsRegion),
			Endpoint:         aws.String(awsEndpoint),
			S3ForcePathStyle: aws.Bool(true),
			Credentials:      credentials.NewStaticCredentials("hau", "hau", ""),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-west-2"))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/MyService/MyApp/Dev/DATABASE_URI"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	value := *param.Parameter.Value
	fmt.Println(value)
	return value, nil
}
