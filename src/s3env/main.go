package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	AWS_ACCESS_KEY_ID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY := os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWS_DEFAULT_REGION := os.Getenv("AWS_DEFAULT_REGION")
	S3_BUCKET_NAME := os.Getenv("S3_BUCKET_NAME")

	creds := credentials.NewStaticCredentials(
		AWS_ACCESS_KEY_ID,
		AWS_SECRET_ACCESS_KEY,
		"",
	)
	sess := session.Must(session.NewSession())
	config := aws.NewConfig().WithRegion(AWS_DEFAULT_REGION).WithCredentials(creds)
	client := s3.New(sess, config)

	result, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(S3_BUCKET_NAME),
	})

	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("%+v\n", result)
	}
}
