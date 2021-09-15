package common

import (
	"os"
	"time"
)

type JobStatus int

const (
	JOB_SUBMITTED JobStatus = iota
	JOB_COMPLETED
	JOB_ERROR
)

var JobStatusString = []string{
	"SUBMITTED",
	"COMPLETED",
	"ERROR",
}

func (m JobStatus) String() string {
	return JobStatusString[m]
}

type VideoRecord struct {
	Username  string    `dynamodbav:"Username"`
	ID        string    `dynamodbav:"ID"`
	CreatedAt time.Time `dynamodbav:"CreatedAt"`
	UpdatedAt time.Time `dynamodbav:"CreatedAt"`
	Duration  string    `dynamodbav:"Duration,omitempty"`
	Width     string    `dynamodbav:"Width,omitempty"`
	Height    string    `dynamodbav:"Height,omitempty"`
	Title     string    `dynamodbav:"Title,omitempty"`
	Key       string    `dynamodbav:"Key,omitempty"`
	JobStatus string    `dynamodbav:"JobStatus"`
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

var (
	Region               = os.Getenv("REGION")
	Stage                = os.Getenv("STAGE")
	VideoTableName       = os.Getenv("VIDEOS_TABLE_NAME")
	MediaConvertURL      = os.Getenv("MEDIACONVERT_URL")
	MediaConvertQueueArn = os.Getenv("MEDIACONVERT_QUEUE_ARN")
	MediaConvertRoleArn  = os.Getenv("MEDIACONVERT_ROLE_ARN")
)
