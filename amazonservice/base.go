package amazonservice

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsSess *session.Session

func InitAws() error {
	var err error
	AccessKeyID := os.Getenv("AWS_KEY_ID")
	SecretAccessKey := os.Getenv("AWS_SECRET")
	MyRegion := os.Getenv("AWS_REGION")
	awsSess, err = session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		return err
	}

	return err
}

func AwsManager() *session.Session {
	return awsSess
}
