package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, e events.CloudWatchEvent) {
	var out bytes.Buffer
	json.Indent(&out, e.Detail, "", "\t")
	log.Print(out.String())
}

func main() {
	lambda.Start(Handler)
}
