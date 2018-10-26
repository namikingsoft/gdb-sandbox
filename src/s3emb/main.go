package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// replaced by sed
	AWS_ACCESS_KEY_ID := "__AWS_ACCESS_KEY_ID__"
	AWS_SECRET_ACCESS_KEY := "__AWS_SECRET_ACCESS_KEY__"
	AWS_DEFAULT_REGION := "__AWS_DEFAULT_REGION__"
	S3_BUCKET_NAME := "__S3_BUCKET_NAME__"

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
