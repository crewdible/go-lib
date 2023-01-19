package amazonservice

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func generateKey(createdAt, folder, key, filePath string) string {
	var fileName string
	lfile := strings.Split(filePath, "/")
	if lfile[len(lfile)-2] == "thumbs" {
		fileName = fmt.Sprintf("thumbs/%s", lfile[len(lfile)-1])
	} else {
		fileName = lfile[len(lfile)-1]
	}
	if key == "" {
		return fmt.Sprintf("%s/%s/%s", createdAt, folder, fileName)
	}

	return fmt.Sprintf("%s/%s/%s/%s", createdAt, folder, key, fileName)
}

func generateCustomKey(folder, key, filePath string) string {
	var fileName string
	lfile := strings.Split(filePath, "/")
	if lfile[len(lfile)-2] == "thumbs" {
		fileName = fmt.Sprintf("thumbs/%s", lfile[len(lfile)-1])
	} else {
		fileName = lfile[len(lfile)-1]
	}

	if key == "" {
		return fmt.Sprintf("%s/%s", folder, fileName)
	}

	return fmt.Sprintf("%s/%s/%s", folder, key, fileName)
}

// func (a AwsSession) UploadFileToS3(bucket, key, acl, cntntDisposition, sSEnc, strgClass, fileName string) error {
func (a AwsSession) UploadFileToS3(bucket, key, acl, filePath, folder, createdAt string) error {
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

// request example UploadCrewFileToS3("ATTACHMENT", "public-read", "./files/pdf/output.pdf", "inv", "202208")
func (a AwsSession) UploadCrewFileToS3(t, acl, filePath, folder, createdAt string) error {
	bkList := getBucketKey(t)
	bucket := bkList[0]
	key := bkList[1]
	err := a.UploadFileToS3(bucket, key, acl, filePath, folder, createdAt)
	if err != nil {
		return err
	}
	lfile := strings.Split(filePath, "/")
	thumbsPath := fmt.Sprintf("./files/thumbs/%s", lfile[len(lfile)-1])
	if _, err := os.Stat(thumbsPath); errors.Is(err, os.ErrNotExist) {
	} else {
		err := a.UploadFileToS3(bucket, key, acl, thumbsPath, folder, createdAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a AwsSession) UploadFileToS3WithReader(bucket, key, acl string, file io.Reader, fileSize int64, filePath, folder, createdAt string) error {
	// open the file for use
	// get the file size and read
	// the file content into a buffer

	buffer := make([]byte, fileSize)
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
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(generateCustomKey(folder, key, filePath)),
		ACL:           aclType,
		Body:          bytes.NewBuffer(buffer),
		ContentLength: *aws.Int64(fileSize),
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

func (a AwsSession) UploadCrewFileToS3WithReader(t, acl string, file io.Reader, fileSize int64, filePath, folder, createdAt string) error {
	bkList := getBucketKey(t)
	bucket := bkList[0]
	key := bkList[1]
	err := a.UploadFileToS3WithReader(bucket, key, acl, file, fileSize, filePath, folder, createdAt)
	if err != nil {
		return err
	}
	lfile := strings.Split(filePath, "/")
	thumbsPath := fmt.Sprintf("./files/thumbs/%s", lfile[len(lfile)-1])
	if _, err := os.Stat(thumbsPath); errors.Is(err, os.ErrNotExist) {
	} else {
		err := a.UploadFileToS3WithReader(bucket, key, acl, file, fileSize, thumbsPath, folder, createdAt)
		if err != nil {
			return err
		}
	}

	return nil
}
