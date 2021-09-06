package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Handler(ctx context.Context, s3Event events.S3Event) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)

	for _, record := range s3Event.Records {
		headObjIn := &s3.HeadObjectInput{
			Bucket: &record.S3.Bucket.Name,
			Key:    &record.S3.Object.Key,
		}

		headObj, err := s3Client.HeadObject(ctx, headObjIn)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(headObj.Metadata)
	}
}

func main() {
	lambda.Start(Handler)
}
