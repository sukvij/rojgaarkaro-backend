package model

import "github.com/spf13/viper"

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	UploadTimeout   int
	BaseURL         string
}

func GetAWSConfig() AWSConfig {
	var awsConfig AWSConfig
	viper.SetConfigFile("config/config.env")
	viper.ReadInConfig()
	awsConfig.AccessKeyID = viper.Get("AWS_ACCESS_KEY_ID").(string)
	awsConfig.Region = viper.Get("AWS_REGION").(string)
	awsConfig.AccessKeySecret = viper.Get("AWS_SECRET_ACCESS_KEY").(string)
	awsConfig.BucketName = viper.Get("BUCKET").(string)
	return awsConfig
}
