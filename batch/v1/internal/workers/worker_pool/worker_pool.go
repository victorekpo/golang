package worker_pool

import (
	"batch-v1/internal/workers/custom_worker"
	"batch-v1/internal/workers/processor"
	"batch-v1/internal/workers/queue"
	"batch-v1/internal/workers/queue_item"
	"fmt"
	"sync"
)

type CustomWorker = custom_worker.CustomWorker
type QueueItem = queue_item.QueueItem
type Queue = queue.Queue

type WorkerPool struct {
	Workers    []*CustomWorker
	NumWorkers int
	Queue      *Queue
	wg         *sync.WaitGroup
}

func NewWorkerPool(numWorkers int, queue *Queue, wg *sync.WaitGroup) *WorkerPool {
	pool := &WorkerPool{
		NumWorkers: numWorkers,
		Queue:      queue,
		wg:         wg,
	}

	// Create processor and share across workers
	customProcessor := processor.NewProcessor("Custom Processor")
	// Create workers
	fmt.Printf("Creating %d workers\n", numWorkers)
	for i := 1; i <= numWorkers; i++ {
		worker := custom_worker.NewCustomWorker(i, pool.Queue, pool.wg, customProcessor)
		pool.Workers = append(pool.Workers, worker)
		go worker.Start()
	}
	return pool
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}
