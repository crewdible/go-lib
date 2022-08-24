package amazonservice

import (
	"bytes"
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (a AwsSession) UploadJSONToS3(acl, filePath, folder, createdAt string) error {
	bkList := getBucketKey("JSON")
	bucket := bkList[0]
	key := bkList[1]

	// open the file for use
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// get the file size and read
	// the file content into a buffer
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	var aclType types.ObjectCannedACL

	switch aclT := acl; aclT {
	case "public-read":
		aclType = types.ObjectCannedACLPublicRead
	default:
		aclType = types.ObjectCannedACLPrivate
	}

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	uploader := manager.NewUploader(a.S3Client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(generateKey(createdAt, folder, key, filePath)),
		ACL:           aclType,
		Body:          bytes.NewReader(buffer),
		ContentLength: *aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
		// ContentDisposition: aws.String(cntntDisposition),
		// ServerSideEncryption: aws.String(sSEnc),
		// StorageClass:         aws.String(strgClass),
	})
	if err != nil {
		return err
	}

	return err
}
