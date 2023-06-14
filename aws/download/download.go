package download

import (
	"fmt"
	"os"

	awsModel "rojgaarkaro-backend/aws/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Download(item string) {
	awsConfig := awsModel.GetAWSConfig()

	// item = "https://rojgaarkaro.s3.ap-south-1.amazonaws.com/users/1.jpg"

	file, err := os.Create(item)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(awsConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsConfig.AccessKeyID, awsConfig.AccessKeySecret, ""),
	})

	downloader := s3manager.NewDownloader(sess)

	fmt.Println(downloader)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(awsConfig.BucketName),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
