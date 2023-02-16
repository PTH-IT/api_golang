package utils

import (
	"PTH-IT/api_golang/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetParameter(name string) (string, error) {
	awsEndpoint := fmt.Sprintf("%s:%s", config.Getconfig().Aws.Host, config.Getconfig().Aws.Port)
	awsRegion := config.Getconfig().Aws.Region
	id := config.Getconfig().Aws.Id
	secret := config.Getconfig().Aws.Secret
	token := config.Getconfig().Aws.Token
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:           aws.String(awsRegion),
			Endpoint:         aws.String(awsEndpoint),
			S3ForcePathStyle: aws.Bool(true),
			Credentials:      credentials.NewStaticCredentials(id, secret, token),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(awsRegion))
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
