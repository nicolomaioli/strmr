package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	mcTypes "github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nicolomaioli/strmr-infra/serverless/video/common"
)

func Handler(ctx context.Context, s3Event events.S3Event) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	dbClient := dynamodb.NewFromConfig(cfg)

	mcClient := mediaconvert.New(mediaconvert.Options{
		Region:      common.Region,
		Credentials: cfg.Credentials,
		EndpointResolver: mediaconvert.EndpointResolverFromURL(
			common.MediaConvertURL,
		),
	})

	for _, record := range s3Event.Records {
		// Fetch video metadata
		headObjIn := &s3.HeadObjectInput{
			Bucket: &record.S3.Bucket.Name,
			Key:    &record.S3.Object.Key,
		}

		headObj, err := s3Client.HeadObject(ctx, headObjIn)
		if err != nil {
			log.Fatal(err)
		}

		// Create a new DynamoDB record for the video
		r := &common.VideoRecord{
			Username: headObj.Metadata["username"],
			VideoID:  headObj.Metadata["id"],
			CreatedAt: fmt.Sprint(
				time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
			),
			Duration: headObj.Metadata["duration"],
			Width:    headObj.Metadata["width"],
			Height:   headObj.Metadata["height"],
			Title:    headObj.Metadata["title"],
			Status:   "SUBMITTED",
		}

		item, err := attributevalue.MarshalMap(r)
		if err != nil {
			log.Fatal(err)
		}

		_, err = dbClient.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: &common.VideoTableName,
			Item:      item,
		})
		if err != nil {
			log.Fatal(err)
		}

		// Start MediaConvert job
		inputS3URI := fmt.Sprintf(
			"s3://%s/%s",
			record.S3.Bucket.Name,
			record.S3.Object.Key,
		)

		outputS3URI := fmt.Sprintf(
			"s3://%s/public/vod/%s/%s",
			record.S3.Bucket.Name,
			headObj.Metadata["username"],
			headObj.Metadata["id"],
		)

		// Values are pulled from the default values assigned when creating a
		// Job in the MediaConvert Console
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

		output, err := mcClient.CreateJob(ctx, createJobInput)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(output)
	}
}

func main() {
	lambda.Start(Handler)
}
