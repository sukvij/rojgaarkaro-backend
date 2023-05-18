package upload

import (
	"mime/multipart"
	"strconv"

	awsModel "rojgaarkaro-backend/aws/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(fileheader *multipart.FileHeader, userId int64) (string, error) {
	awsConfig := awsModel.GetAWSConfig()
	s3Config := &aws.Config{
		Region:      aws.String(awsConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsConfig.AccessKeyID, awsConfig.AccessKeySecret, ""),
	}
	s3Session, _ := session.NewSession(s3Config)

	uploader := s3manager.NewUploader(s3Session)
	file, _ := fileheader.Open()
	input := &s3manager.UploadInput{
		Bucket: aws.String(awsConfig.BucketName),                          // bucket's name
		Key:    aws.String("users/" + strconv.Itoa(int(userId)) + ".jpg"), // files destination location
		Body:   file,                                                      // content type
	}
	output, err := uploader.Upload(input)
	if err != nil {
		return "", err
	}
	location := output.Location
	return location, nil
}
