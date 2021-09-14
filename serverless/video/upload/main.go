package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	mcTypes "github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nicolomaioli/strmr-infra/serverless/video/common"
)

func getObject(ctx context.Context, cfg aws.Config, r *events.S3EventRecord) (*s3.HeadObjectOutput, error) {
	client := s3.NewFromConfig(cfg)

	headObjIn := &s3.HeadObjectInput{
		Bucket: &r.S3.Bucket.Name,
		Key:    &r.S3.Object.Key,
	}

	return client.HeadObject(ctx, headObjIn)
}

func submitJob(ctx context.Context, cfg aws.Config, r *events.S3EventRecord, obj *s3.HeadObjectOutput) common.JobStatus {
	client := mediaconvert.New(mediaconvert.Options{
		Region:      common.Region,
		Credentials: cfg.Credentials,
		EndpointResolver: mediaconvert.EndpointResolverFromURL(
			common.MediaConvertURL,
		),
	})

	var jobStatus = common.JOB_SUBMITTED

	inputS3URI := fmt.Sprintf(
		"s3://%s/%s",
		r.S3.Bucket.Name,
		r.S3.Object.Key,
	)

	outputS3URI := fmt.Sprintf(
		"s3://%s/public/vod/%s/%s",
		r.S3.Bucket.Name,
		obj.Metadata["username"],
		obj.Metadata["id"],
	)

	// Beware the monster, this took a minute to figure out
	createJobInput := &mediaconvert.CreateJobInput{
		Role:  &common.MediaConvertRoleArn,
		Queue: &common.MediaConvertQueueArn,
		Settings: &mcTypes.JobSettings{
			Inputs: []mcTypes.Input{
				{
					FileInput: &inputS3URI,
					VideoSelector: &mcTypes.VideoSelector{
						ColorSpace: "FOLLOW",
					},
					AudioSelectors: map[string]mcTypes.AudioSelector{
						"Audio Selector 1": {
							Offset:           0,
							DefaultSelection: "DEFAULT",
							SelectorType:     "LANGUAGE_CODE",
							ProgramSelection: 1,
							LanguageCode:     "ENM",
						},
					},
				},
			},
			OutputGroups: []mcTypes.OutputGroup{
				{
					Outputs: []mcTypes.Output{
						{
							VideoDescription: &mcTypes.VideoDescription{
								CodecSettings: &mcTypes.VideoCodecSettings{
									Codec: mcTypes.VideoCodecH264,
									H264Settings: &mcTypes.H264Settings{
										RateControlMode: mcTypes.H264RateControlModeQvbr,
										MaxBitrate:      5000000,
									},
								},
							},
							AudioDescriptions: []mcTypes.AudioDescription{
								{
									CodecSettings: &mcTypes.AudioCodecSettings{
										Codec: mcTypes.AudioCodecAac,
										AacSettings: &mcTypes.AacSettings{
											CodingMode: mcTypes.AacCodingModeCodingMode10,
											SampleRate: 48000,
											Bitrate:    96000,
										},
									},
								},
							},
							ContainerSettings: &mcTypes.ContainerSettings{
								Container: mcTypes.ContainerTypeMpd,
							},
						},
					},
					OutputGroupSettings: &mcTypes.OutputGroupSettings{
						Type: "DASH_ISO_GROUP_SETTINGS",
						DashIsoGroupSettings: &mcTypes.DashIsoGroupSettings{
							Destination:    &outputS3URI,
							SegmentLength:  30,
							FragmentLength: 2,
						},
					},
				},
			},
		},
	}

	_, err := client.CreateJob(ctx, createJobInput)
	if err != nil {
		jobStatus = common.JOB_ERROR
	}

	return jobStatus
}

func putRecord(ctx context.Context, cfg aws.Config, obj *s3.HeadObjectOutput, status common.JobStatus) error {
	client := dynamodb.NewFromConfig(cfg)

	r := &common.VideoRecord{
		Username: obj.Metadata["username"],
		VideoID:  obj.Metadata["id"],
		CreatedAt: fmt.Sprint(
			time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		),
		Duration: obj.Metadata["duration"],
		Width:    obj.Metadata["width"],
		Height:   obj.Metadata["height"],
		Title:    obj.Metadata["title"],
		Status:   status.String(),
	}

	item, err := attributevalue.MarshalMap(r)
	if err != nil {
		return err
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &common.VideoTableName,
		Item:      item,
	})
	if err != nil {
		return err
	}

	return nil
}

func Handler(ctx context.Context, s3Event events.S3Event) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range s3Event.Records {
		// Fetch video object
		obj, err := getObject(ctx, cfg, &r)
		if err != nil {
			log.Fatal(err)
		}

		// Submit MediaConvert job
		status := submitJob(ctx, cfg, &r, obj)

		// Create a new DynamoDB record for the video
		err = putRecord(ctx, cfg, obj, status)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	lambda.Start(Handler)
}
