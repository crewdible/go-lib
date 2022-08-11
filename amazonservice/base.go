package amazonservice

import (
	"context"
	"os"

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

func getBucketKey(t string) []string {
	envName := os.Getenv("ENV_NAME")
	if envName == "PRODUCTION" {
		switch bk := t; bk {
		case "SKU":
			return []string{"crewdible-pub", "sku"}
		case "USER":
			return []string{"crewdible-pub", "user"}
		case "GALLERIES":
			return []string{"crewdible-pub", "galleries"}
		case "FACILITY":
			return []string{"crewdible-pub", "facilities"}
		case "MARKETPLACE":
			return []string{"crewdible-pub", "mp"}
		case "PACKAGING":
			return []string{"crewdible-pub", "packaging"}
		case "LOGISTIC":
			return []string{"crewdible-pub", "logistic"}
		case "LOGISTICOLD":
			return []string{"crewdible-pub", "logistic"}
		case "ATTACHMENT":
			return []string{"crewdible-outbound", "attachment"}
		case "JSON":
			return []string{"crewdible-outbound", ""}
		default:
			return []string{"crewdible-pub", "test"}
		}
	} else {
		switch bk := t; bk {
		case "SKU":
			return []string{"crewdible-sandbox-pub", "sku"}
		case "USER":
			return []string{"crewdible-sandbox-pub", "user"}
		case "GALLERIES":
			return []string{"crewdible-sandbox-pub", "galleries"}
		case "FACILITY":
			return []string{"crewdible-sandbox-pub", "facilities"}
		case "MARKETPLACE":
			return []string{"crewdible-sandbox-pub", "mp"}
		case "PACKAGING":
			return []string{"crewdible-sandbox-pub", "packaging"}
		case "LOGISTIC":
			return []string{"crewdible-sandbox-pub", "logistic"}
		case "LOGISTICOLD":
			return []string{"crewdible-sandbox-pub", "logistic"}
		case "ATTACHMENT":
			return []string{"crewdible-sandbox-outbound", "attachment"}
		case "JSON":
			return []string{"crewdible-sandbox-outbound", ""}
		default:
			return []string{"crewdible-sandbox-pub", "test"}
		}
	}
}
