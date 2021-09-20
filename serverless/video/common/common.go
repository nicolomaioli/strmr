package common

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

var (
	Region               = os.Getenv("REGION")
	Stage                = os.Getenv("STAGE")
	OutputBucketName     = os.Getenv("OUTPUT_BUCKET_NAME")
	VideoTableName       = os.Getenv("VIDEOS_TABLE_NAME")
	MediaConvertURL      = os.Getenv("MEDIACONVERT_URL")
	MediaConvertQueueArn = os.Getenv("MEDIACONVERT_QUEUE_ARN")
	MediaConvertRoleArn  = os.Getenv("MEDIACONVERT_ROLE_ARN")
	ServeVideoURL        = os.Getenv("SERVE_VIDEO_URL")

	LambdaHTTPHeaders = map[string]string{
		"Content-Type": "application/json",
	}
)

type JobStatus int

const (
	JOB_SUBMITTED JobStatus = iota
	JOB_COMPLETED
	JOB_ERROR
	JOB_TEST
)

var JobStatusString = []string{
	"SUBMITTED",
	"COMPLETED",
	"ERROR",
	"TEST",
}

func (m JobStatus) String() string {
	return JobStatusString[m]
}

type LambdaHTTPErrorBody struct {
	Message string `json:"message"`
}

type VideoRecord struct {
	Username  string
	ID        string
	Duration  float64
	Width     int
	Height    int
	Title     string
	JobStatus string
	CreatedAt time.Time `dynamodbav:",unixtime"`
	UpdatedAt time.Time `dynamodbav:",unixtime"`
	Path      string    `dynamodbav:",omitempty"`
}

type MediaConvertEventDetail struct {
	Timestamp          int               `json:"timestamp"`
	AccountId          string            `json:"accountId"`
	Queue              string            `json:"queue"`
	JobId              string            `json:"jobId"`
	Status             string            `json:"status"`
	UserMetadata       map[string]string `json:"userMetadata"`
	OutputGroupDetails []struct {
		OutputDetails []struct {
			OutputFilePaths []string `json:"outputFilePaths"`
			DurationInMs    int      `json:"durationInMs"`
			VideoDetails    struct {
				WidthInPx  int `json:"widthInPx"`
				HeightInPx int `json:"heightInPx"`
			} `json:"videoDetails"`
		} `json:"outputDetails"`
		PlaylistFilePaths []string `json:"playlistFilePaths"`
		Type              string   `json:"type"`
	} `json:"outputGroupDetails"`
}

func HandleLambdaHTTPError(ctx context.Context, err error, status int) *events.APIGatewayV2HTTPResponse {
	log.Print(err)

	if status > 499 {
		err = errors.New("internal server error")
	}

	e := LambdaHTTPErrorBody{
		Message: err.Error(),
	}

	b, _ := json.Marshal(e)

	return &events.APIGatewayV2HTTPResponse{
		Headers:    LambdaHTTPHeaders,
		Body:       string(b),
		StatusCode: status,
	}
}
