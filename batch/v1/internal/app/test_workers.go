package app

import (
	"batch-v1/internal/config"
	"batch-v1/internal/workers/queue"
	"batch-v1/internal/workers/queue_item"
	"batch-v1/internal/workers/worker_pool"
	"fmt"
	"sync"
)

func TestWorkers() {
	var wg sync.WaitGroup
	numWorkers := config.NumberOfWorkers
	maxBuffer := config.MaxQueueChannelBuffer
	q := queue.NewQueue(maxBuffer)
	pool := worker_pool.NewWorkerPool(numWorkers, q, &wg)

	item1 := queue_item.NewQueueItem("data1")
	item2 := queue_item.NewQueueItem("data2")
	item3 := queue_item.NewQueueItem("data3")
	item4 := queue_item.NewQueueItem("data4")
	item5 := queue_item.NewQueueItem("data5")

	q.Add(item1)
	q.Add(item2)
	q.Add(item3)
	q.Add(item4)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	q.Add(item5)
	fmt.Println(q)

	pool.Wait()
	fmt.Println(q)

	// If items in the DLQ try to reprocess them
	q.ProcessAllDLQ()
	pool.Wait()
	fmt.Println(q)
}
