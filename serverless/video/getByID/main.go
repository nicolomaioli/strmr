package main

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/nicolomaioli/strmr/serverless/video/common"
)

func getByID(ctx context.Context, cfg aws.Config, id string) (*common.VideoRecord, error) {
	client := dynamodb.NewFromConfig(cfg)

	key := map[string]string{
		"ID": id,
	}

	pk, err := attributevalue.MarshalMap(key)
	if err != nil {
		return nil, err
	}

	proj := expression.NamesList(
		expression.Name("ID"),
		expression.Name("Username"),
		expression.Name("Title"),
		expression.Name("Duration"),
		expression.Name("Width"),
		expression.Name("Height"),
		expression.Name("Path"),
		expression.Name("JobStatus"),
		expression.Name("CreatedAt"),
		expression.Name("UpdatedAt"),
	)

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName:                &common.VideoTableName,
		Key:                      pk,
		ExpressionAttributeNames: expr.Names(),
		ProjectionExpression:     expr.Projection(),
	}

	result, err := client.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	r := common.VideoRecord{}
	if err := attributevalue.UnmarshalMap(result.Item, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		e := common.HandleLambdaHTTPError(err, 500)
		return *e, nil
	}

	if _, ok := req.PathParameters["id"]; !ok {
		err := errors.New("request: missing require path parameter \"id\"")
		e := common.HandleLambdaHTTPError(err, 400)
		return *e, nil
	}

	r, err := getByID(ctx, cfg, req.PathParameters["id"])
	if err != nil {
		e := common.HandleLambdaHTTPError(err, 500)
		return *e, nil
	}

	if r.ID == "" {
		err := errors.New("request: video not found")
		e := common.HandleLambdaHTTPError(err, 404)
		return *e, nil
	}

	b, err := json.Marshal(r)
	if err != nil {
		e := common.HandleLambdaHTTPError(err, 500)
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
