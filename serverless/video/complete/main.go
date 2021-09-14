package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nicolomaioli/strmr-infra/serverless/video/common"
)

func Handler(ctx context.Context, e events.CloudWatchEvent) {
	var d common.MediaConvertEventDetail
	err := json.Unmarshal(e.Detail, &d)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Completed video: %s", d.UserMetadata["id"])
}

func main() {
	lambda.Start(Handler)
}
