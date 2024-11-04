package custom_worker

import (
	"batch-v1/internal/config"
	"batch-v1/internal/processor"
	"batch-v1/internal/queue_item"
	"batch-v1/internal/utils"
	"fmt"
	"sync"
	"time"
)

type WorkerStatus string
type QueueItem = queue_item.QueueItem

const (
	WorkerIdle WorkerStatus = "idle"
	WorkerBusy WorkerStatus = "busy"
)

type CustomWorker struct {
	ID         int
	Status     WorkerStatus
	Queue      chan *QueueItem
	DLQ        chan *QueueItem
	mu         sync.Mutex
	wg         *sync.WaitGroup
	processor  *processor.Processor
	MaxRetries int
}

func NewCustomWorker(id int, queue chan *QueueItem, dlq chan *QueueItem, wg *sync.WaitGroup, processor *processor.Processor) *CustomWorker {
	return &CustomWorker{
		ID:         id,
		Status:     WorkerIdle,
		Queue:      queue,
		DLQ:        dlq,
		wg:         wg,
		processor:  processor,
		MaxRetries: config.MaxRetries,
	}
}

func (w *CustomWorker) Start() {
	for item := range w.Queue {
		w.SetStatus(WorkerBusy)
		// Process Item with retry handler
		w.RetryHandler(func() error {
			fmt.Printf("Worker #%v processing item %s\n", w.ID, item.Data)
			return w.processor.ProcessItem(item)
		}, item)
		w.SetStatus(WorkerIdle)
		w.wg.Done() // Decrement the WaitGroup counter after processing each item
	}
}

func (w *CustomWorker) RetryHandler(action func() error, item *QueueItem) {
	for retries := 0; retries < w.MaxRetries; retries++ {
		err := action()
		if err == nil {
			return
		}
		fmt.Printf("Error processing item %s, attempt %d: %s\n", item.Data, retries+1, err)
		time.Sleep(utils.ExponentialBackoff(retries))
	}
	fmt.Printf("Failed to process item %s after %d attempts, sending to DLQ\n", item.Data, w.MaxRetries)
	w.DLQ <- item
}

func (w *CustomWorker) SetStatus(status WorkerStatus) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Status = status
	fmt.Printf("Worker #%v set to %v\n", w.ID, w.Status)
}

func (w *CustomWorker) IsIdle() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.Status == WorkerIdle
}
