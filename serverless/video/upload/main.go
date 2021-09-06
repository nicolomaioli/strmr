package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nicolomaioli/strmr-infra/serverless/video/common"
)

func Handler(ctx context.Context, s3Event events.S3Event) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	dbClient := dynamodb.NewFromConfig(cfg)

	for _, record := range s3Event.Records {
		headObjIn := &s3.HeadObjectInput{
			Bucket: &record.S3.Bucket.Name,
			Key:    &record.S3.Object.Key,
		}

		headObj, err := s3Client.HeadObject(ctx, headObjIn)
		if err != nil {
			log.Fatal(err)
		}

		r := &common.VideoRecord{
			Username: headObj.Metadata["username"],
			VideoID:  headObj.Metadata["id"],
			CreatedAt: fmt.Sprint(
				time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
			),
			Duration: headObj.Metadata["duration"],
			Width:    headObj.Metadata["width"],
			Height:   headObj.Metadata["height"],
			Title:    headObj.Metadata["title"],
			Key:      record.S3.Object.Key,
			FileType: "VIDEO",
		}

		item, err := attributevalue.MarshalMap(r)
		if err != nil {
			log.Fatal(err)
		}

		output, err := dbClient.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: &common.VideoTableName,
			Item:      item,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Print(output.ResultMetadata)
	}
}

func main() {
	lambda.Start(Handler)
}
