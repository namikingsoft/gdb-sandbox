package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"app/src/lib"
)

func main() {
	result, err := s3.New(
		session.Must(session.NewSession()),
		aws.NewConfig().
			WithRegion(lib.Decrypt("b106da887e41cbc6f0848a73d8b0")).
			WithCredentials(
				credentials.NewStaticCredentials(
					lib.Decrypt("913dbea75b67f9efd1b5b246b6b3e0db1bf37e25"),
					lib.Decrypt("810fa580454190ddf8d5924985cb9f9a5994f54fc940b9c9c7fd9f33a6c5d13504194ec134c59a94"),
					"",
				),
			),
	).ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(lib.Decrypt("b7129595705ddbccfa9d")),
	})

	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("%+v\n", result)
	}
}
