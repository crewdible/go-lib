package amazonservice

import (
	"context"
	"fmt"
	"os"
	"strings"

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

func GetAwsUrl(t, folder, fileName string, isThumb bool) string {
	var tmb string
	bkList := getBucketKey(t)
	bucket := bkList[0]
	key := bkList[1]

	if key != "" {
		key = fmt.Sprintf("%s/", key)
	}

	if isThumb {
		tmb = "thumbs/"
	}

	return fmt.Sprintf("https://%s.s3.ap-southeast-1.amazonaws.com/%s%s%s%s", bucket, folder, key, tmb, fileName)
}

// default t => "ATTACHMENT", default isThumb => false
func GetAttachments(t, files, folder, createdAt string, isThumb bool) string {
	var awsResults []string
	awsFiles := strings.Split(files, ";")
	for _, file := range awsFiles {
		atc := strings.Contains(file, "attachment")
		if atc {
			awsResults = append(awsResults, file)
		} else {
			if isThumb && len(file) > 4 && file[len(file)-4:] == ".pdf" {
				continue
			} else {
				folderAws := fmt.Sprintf("%s/%s/", createdAt, folder)
				awsUrl := GetAwsUrl(t, folderAws, file, isThumb)
				awsResults = append(awsResults, awsUrl)
			}
		}
	}

	return strings.Join(awsResults, ";")
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
		case "EXCEL":
			return []string{"crewdible-pub", "excel"}
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
		case "EXCEL":
			return []string{"crewdible-sandbox-pub", "excel"}
		default:
			return []string{"crewdible-sandbox-pub", "test"}
		}
	}
}

var productionBucketNames = map[string]string{
	"pub":      "crewdible-pub",
	"outbound": "crewdible-outbound",
}

var nonProductionBucketNames = map[string]string{
	"pub":      "crewdible-sandbox-pub",
	"outbound": "crewdible-sandbox-outbound",
}

func getBucketName(key string) string {
	envName := os.Getenv("ENV_NAME")
	if envName == "PRODUCTION" {
		name, exist := productionBucketNames[key]
		if exist {
			return name
		}
		return "crewdible-" + name
	}

	name, exist := nonProductionBucketNames[key]
	if exist {
		return name
	}

	return "crewdible-sandbox-" + name
}
