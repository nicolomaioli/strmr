package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
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
		common.OutputBucketName,
		obj.Metadata["username"],
		obj.Metadata["id"],
	)

	// Beware the monster, this took a minute to figure out
	createJobInput := &mediaconvert.CreateJobInput{
		Role:  &common.MediaConvertRoleArn,
		Queue: &common.MediaConvertQueueArn,
		UserMetadata: map[string]string{
			"id":       obj.Metadata["id"],
			"username": obj.Metadata["username"],
		},
		Settings: &types.JobSettings{
			Inputs: []types.Input{
				{
					FileInput: &inputS3URI,
					VideoSelector: &types.VideoSelector{
						ColorSpace: "FOLLOW",
					},
					AudioSelectors: map[string]types.AudioSelector{
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
			OutputGroups: []types.OutputGroup{
				{
					Outputs: []types.Output{
						{
							VideoDescription: &types.VideoDescription{
								CodecSettings: &types.VideoCodecSettings{
									Codec: types.VideoCodecH264,
									H264Settings: &types.H264Settings{
										RateControlMode: types.H264RateControlModeQvbr,
										MaxBitrate:      5000000,
									},
								},
							},
							AudioDescriptions: []types.AudioDescription{
								{
									CodecSettings: &types.AudioCodecSettings{
										Codec: types.AudioCodecAac,
										AacSettings: &types.AacSettings{
											CodingMode: types.AacCodingModeCodingMode10,
											SampleRate: 48000,
											Bitrate:    96000,
										},
									},
								},
							},
							ContainerSettings: &types.ContainerSettings{
								Container: types.ContainerTypeMpd,
							},
						},
					},
					OutputGroupSettings: &types.OutputGroupSettings{
						Type: "DASH_ISO_GROUP_SETTINGS",
						DashIsoGroupSettings: &types.DashIsoGroupSettings{
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
	timestamp := time.Now()

	duration, err := strconv.ParseFloat(obj.Metadata["duration"], 64)
	if err != nil {
		return err
	}

	width, err := strconv.Atoi(obj.Metadata["width"])
	if err != nil {
		return err
	}

	height, err := strconv.Atoi(obj.Metadata["height"])
	if err != nil {
		return err
	}

	r := &common.VideoRecord{
		Username:  obj.Metadata["username"],
		ID:        obj.Metadata["id"],
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
		Duration:  duration,
		Width:     width,
		Height:    height,
		Title:     obj.Metadata["title"],
		JobStatus: status.String(),
	}

	item, err := attributevalue.MarshalMap(r)
	if err != nil {
		return err
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &common.VideoTableName,
		Item:      item,
	})

	return err
}

func Handler(ctx context.Context, e events.S3Event) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range e.Records {
		// Fetch video object
		obj, err := getObject(ctx, cfg, &r)
		if err != nil {
			log.Fatal(err)
		}

		// Submit MediaConvert job, unless it's a test
		var status common.JobStatus

		if _, ok := obj.Metadata["test"]; ok {
			log.Print("testing lambda, skipping conversion")
			status = common.JOB_TEST
		} else {
			status = submitJob(ctx, cfg, &r, obj)
		}

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
