package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/nicolomaioli/strmr/serverless/video/common"
)

func queryByJobStatus(ctx context.Context, cfg aws.Config, jobStatus common.JobStatus) (*[]common.VideoRecord, error) {
	client := dynamodb.NewFromConfig(cfg)

	keyCond := expression.
		Key("JobStatus").
		Equal(expression.Value(jobStatus.String()))

	expr, err := expression.
		NewBuilder().
		WithKeyCondition(keyCond).
		Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 &common.VideoTableName,
		IndexName:                 aws.String("JobStatusIndex"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	res, err := client.Query(ctx, input)
	if err != nil {
		return nil, err
	}

	videos := []common.VideoRecord{}
	if err := attributevalue.UnmarshalListOfMaps(res.Items, &videos); err != nil {
		return nil, err
	}

	return &videos, nil
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		e := common.HandleLambdaHTTPError(ctx, err, 500)
		return *e, nil
	}

	r, err := queryByJobStatus(ctx, cfg, common.JOB_COMPLETED)
	if err != nil {
		e := common.HandleLambdaHTTPError(ctx, err, 500)
		return *e, nil
	}

	// I'm happy for the client to deal with an empty array
	b, err := json.Marshal(r)
	if err != nil {
		e := common.HandleLambdaHTTPError(ctx, err, 500)
		return *e, nil
	}

	res := events.APIGatewayV2HTTPResponse{
		Headers:    common.LambdaHTTPHeaders,
		Body:       string(b),
		StatusCode: 200,
	}
	return res, nil
}

func main() {
	lambda.Start(Handler)
}
