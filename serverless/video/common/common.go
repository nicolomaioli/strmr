package common

import "os"

type VideoRecord struct {
	Username  string `dynamodbav:"Username"`
	VideoID   string `dynamodbav:"VideoID"`
	CreatedAt string `dynamodbav:"CreatedAt"`
	Duration  string `dynamodbav:"Duration,omitempty"`
	Width     string `dynamodbav:"Width,omitempty"`
	Height    string `dynamodbav:"Height,omitempty"`
	Title     string `dynamodbav:"Title,omitempty"`
	Key       string `dynamodbav:"Key"`
	FileType  string `dynamodbav:"FileType"`
}

var (
	VideoTableName = os.Getenv("VIDEOS_TABLE_NAME")
)
