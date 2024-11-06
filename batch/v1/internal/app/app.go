package app

import (
	appConfig "batch-v1/internal/config"
	"batch-v1/internal/server"
	"batch-v1/internal/services/sqs"
	"batch-v1/internal/workers/queue"
	"batch-v1/internal/workers/worker_pool"
	"fmt"
	"sync"
)

func App() {
	fmt.Println("Welcome to my App")
	// Create SQS client
	sqsClient, queueUrl := sqs.CreateSQSClient()

	// Initialize worker pool & internal queue
	var wg sync.WaitGroup
	numWorkers := appConfig.NumberOfWorkers
	maxBuffer := appConfig.MaxQueueChannelBuffer
	q := queue.NewQueue(maxBuffer)
	worker_pool.NewWorkerPool(numWorkers, q, &wg)

	// Start polling SQS messages to add to internal queue
	go sqs.PollSQS(sqsClient, queueUrl, q)
	server.Server()
}

// Running PollSQS as a goroutine is similar to asynchronous code in other programming languages.
// It allows the function to execute concurrently with other code, enabling the server to handle
// multiple tasks at the same time without blocking. This is useful for long-running operations like
// polling a message queue while still being able to respond to HTTP requests.

// In Go, if you want to run a function concurrently without blocking the execution of other code, you
// can run it as a goroutine. This is useful for tasks that can be performed in the background or
// concurrently with other operations, such as handling multiple client requests, polling a message queue,
// or performing I/O operations.
