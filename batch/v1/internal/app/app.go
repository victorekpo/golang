package app

import (
	"batch-v1/internal/queue"
	"batch-v1/internal/queue_item"
	"batch-v1/internal/worker_pool"
	"fmt"
	"sync"
)

func App() {
	var wg sync.WaitGroup
	var numWorkers = 5
	q := queue.NewQueue(numWorkers, &wg)
	pool := worker_pool.NewWorkerPool(numWorkers, q)

	item1 := queue_item.NewQueueItem("data1")
	item2 := queue_item.NewQueueItem("error")
	item3 := queue_item.NewQueueItem("data3")
	item4 := queue_item.NewQueueItem("data4")
	item5 := queue_item.NewQueueItem("data5")

	q.Add(item1)
	q.Add(item2)
	q.Add(item3)
	q.Add(item4)
	q.Add(item5)
	fmt.Println(q)

	pool.Wait()
	fmt.Println(q)
}
