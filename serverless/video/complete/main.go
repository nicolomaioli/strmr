package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/nicolomaioli/strmr-infra/serverless/video/common"
)

func updateRecord(ctx context.Context, cfg aws.Config, d *common.MediaConvertEventDetail) error {
	client := dynamodb.NewFromConfig(cfg)

	key := map[string]string{
		"ID": d.UserMetadata["id"],
	}

	pk, err := attributevalue.MarshalMap(key)
	if err != nil {
		log.Fatal(err)
	}

	upd := expression.
		Set(expression.Name("JobStatus"), expression.Value("COMPLETED")).
		Set(expression.Name("Path"), expression.Value(d.OutputGroupDetails[0].PlaylistFilePaths[0])).
		Set(expression.Name("UpdatedAt"), expression.Value(time.Now().Unix()))

	expr, err := expression.NewBuilder().WithUpdate(upd).Build()
	if err != nil {
		log.Fatal(err)
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 &common.VideoTableName,
		Key:                       pk,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = client.UpdateItem(ctx, input)
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
