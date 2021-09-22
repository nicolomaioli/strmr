package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nicolomaioli/strmr/serverless/video/common"
)

func getPosterFrame(ctx context.Context, cfg aws.Config, username, id string) string {
	// Check if .0000001.jpg is present or return .0000000.jpg, which is always
	// present (first frame of video)
	client := s3.NewFromConfig(cfg)

	ext := "0000001.jpg"
	key := fmt.Sprintf("public/%s/%s.%s", username, id, ext)

	headObjIn := &s3.HeadObjectInput{
		Bucket: &common.OutputBucketName,
		Key:    &key,
	}

	_, err := client.HeadObject(ctx, headObjIn)
	if err != nil {
		return "0000000.jpg"
	}

	return ext
}

func updateRecord(ctx context.Context, cfg aws.Config, username, id string) error {
	client := dynamodb.NewFromConfig(cfg)

	basePath := fmt.Sprintf(
		"https://%s/public/%s/%s",
		common.ServeVideoURL,
		username,
		id,
	)

	path := fmt.Sprintf(
		"%s.mpd",
		basePath,
	)

	posterFrame := fmt.Sprintf(
		"%s.%s",
		basePath,
		getPosterFrame(ctx, cfg, username, id),
	)

	key := map[string]string{
		"ID": id,
	}

	pk, err := attributevalue.MarshalMap(key)
	if err != nil {
		log.Fatal(err)
	}

	upd := expression.
		Set(expression.Name("JobStatus"), expression.Value("COMPLETED")).
		Set(expression.Name("Path"), expression.Value(path)).
		Set(expression.Name("PosterFrame"), expression.Value(posterFrame)).
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

	err = updateRecord(
		ctx,
		cfg,
		d.UserMetadata["username"],
		d.UserMetadata["id"],
	)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lambda.Start(Handler)
}
