package sqs

import (
	"batch-v1/internal/workers/queue"
	"batch-v1/internal/workers/queue_item"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
)

func CreateSQSClient() (*sqs.Client, string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == sqs.ServiceID {
					return aws.Endpoint{
						URL: "http://localhost:9324",
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			}),
		),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)
	queueURL := "http://localhost:9324/queue/my-queue"

	return sqsClient, queueURL
}

func PollSQS(client *sqs.Client, queueURL string, q *queue.Queue) {
	for {
		output, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            &queueURL,
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     20,
		})
		if err != nil {
			log.Printf("failed to receive messages, %v", err)
			continue
		}

		for _, message := range output.Messages {
			item := queue_item.NewQueueItem(*message.Body)
			q.Add(item)

			// Delete message from SQS
			_, err := client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
				QueueUrl:      &queueURL,
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				log.Printf("failed to delete message, %v", err)
			}
		}
	}
}

// http://localhost:9324/queue/my-queue?Action=SendMessage&MessageBody=TestMessage1
