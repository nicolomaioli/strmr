package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nicolomaioli/strmr-infra/serverless/video/common"
)

func updateRecord(ctx context.Context, cfg aws.Config, d *common.MediaConvertEventDetail) error {
	client := dynamodb.NewFromConfig(cfg)

	in := &dynamodb.UpdateItemInput{
		TableName: &common.VideoTableName,
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: d.UserMetadata["id"],
			},
		},
		UpdateExpression: aws.String("SET Status = :status, Path = :path"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":status": &types.AttributeValueMemberS{
				Value: "COMPLETED",
			},
			":path": &types.AttributeValueMemberS{
				Value: d.OutputGroupDetails[0].PlaylistFilePaths[0],
			},
		},
	}

	_, err := client.UpdateItem(ctx, in)
	return err
}

func Handler(ctx context.Context, e events.CloudWatchEvent) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var d common.MediaConvertEventDetail
	err = json.Unmarshal(e.Detail, &d)
	if err != nil {
		log.Fatal(err)
	}

	err = updateRecord(ctx, cfg, &d)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lambda.Start(Handler)
}
