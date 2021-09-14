package common

import "os"

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
	Username  string `dynamodbav:"Username"`
	VideoID   string `dynamodbav:"VideoID"`
	CreatedAt string `dynamodbav:"CreatedAt"`
	Duration  string `dynamodbav:"Duration,omitempty"`
	Width     string `dynamodbav:"Width,omitempty"`
	Height    string `dynamodbav:"Height,omitempty"`
	Title     string `dynamodbav:"Title,omitempty"`
	Key       string `dynamodbav:"Key,omitempty"`
	Status    string `dynamodbav:"Status"`
}

var (
	Region               = os.Getenv("REGION")
	Stage                = os.Getenv("STAGE")
	VideoTableName       = os.Getenv("VIDEOS_TABLE_NAME")
	MediaConvertURL      = os.Getenv("MEDIACONVERT_URL")
	MediaConvertQueueArn = os.Getenv("MEDIACONVERT_QUEUE_ARN")
	MediaConvertRoleArn  = os.Getenv("MEDIACONVERT_ROLE_ARN")
)
