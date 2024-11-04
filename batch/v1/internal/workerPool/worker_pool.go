package workerPool

import (
	"batch-v1/internal/customWorker"
	"batch-v1/internal/processor"
	"batch-v1/internal/queue"
	"batch-v1/internal/queueItem"
	"fmt"
	"sync"
)

type CustomWorker = customWorker.CustomWorker
type QueueItem = queueItem.QueueItem
type Queue = queue.Queue

type WorkerPool struct {
	Workers    []*CustomWorker
	NumWorkers int
	Queue      chan *QueueItem
	DLQ        chan *QueueItem
	wg         *sync.WaitGroup
}

func NewWorkerPool(numWorkers int, queue *Queue) *WorkerPool {
	pool := &WorkerPool{
		NumWorkers: numWorkers,
		Queue:      queue.Queue,
		DLQ:        queue.DLQ,
		wg:         queue.Wg,
	}

	// Create processor and share across workers
	customProcessor := processor.NewProcessor("Custom Processor")
	// Create workers
	fmt.Printf("Creating %d workers\n", numWorkers)
	for i := 0; i < numWorkers; i++ {
		worker := customWorker.NewCustomWorker(i, pool.Queue, pool.DLQ, pool.wg, customProcessor)
		pool.Workers = append(pool.Workers, worker)
		go worker.Start()
	}
	return pool
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}
