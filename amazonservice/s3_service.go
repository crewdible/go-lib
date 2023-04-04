package amazonservice

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

	var res string

	if folder != "" {
		res += folder + "/"
	}

	if key != "" {
		res += key + "/"
	}

	return res + fileName
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

func (a AwsSession) UploadFileToS3WithReader(payload UploadWithReaderPayload) error {

	buffer, err := ioutil.ReadAll(payload.File)
	if err != nil {
		return err
	}

	var aclType types.ObjectCannedACL
	switch aclT := payload.Access; aclT {
	case "public-read":
		aclType = types.ObjectCannedACLPublicRead
	default:
		aclType = types.ObjectCannedACLPrivate
	}

	key := strings.Trim(payload.Folder, "/") + "/" + strings.Trim(payload.FileName, "/")
	uploader := manager.NewUploader(a.S3Client)
	output, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(getBucketName(payload.Bucket)),
		Key:           aws.String(key),
		ACL:           aclType,
		Body:          bytes.NewBuffer(buffer),
		ContentLength: *aws.Int64(int64(len(buffer))),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	if err != nil {
		return err
	}

	fmt.Println(output.Location, output.Key)

	return err
}

type UploadWithReaderPayload struct {
	Bucket   string
	File     io.Reader
	FileName string
	Folder   string
	Access   string
}

func (a AwsSession) UploadCrewFileToS3WithReader(payload UploadWithReaderPayload) error {
	return a.UploadFileToS3WithReader(payload)
}
