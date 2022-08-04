package amazonservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsSession struct {
	Config   aws.Config
	S3Client *s3.Client
}

var awsSess AwsSession

func InitAws() error {
	var err error
	awsSess.Config, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	awsSess.S3Client = s3.NewFromConfig(awsSess.Config)

	return err
}

func AwsManager() AwsSession {
	return awsSess
}
