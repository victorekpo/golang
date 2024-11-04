package app

import (
	"batch-v1/internal/queue"
	"batch-v1/internal/queueItem"
	"batch-v1/internal/workerPool"
	"fmt"
	"sync"
)

func App() {
	var wg sync.WaitGroup
	var numWorkers = 5
	q := queue.NewQueue(numWorkers, &wg)
	pool := workerPool.NewWorkerPool(numWorkers, q)

	item1 := queueItem.NewQueueItem("data1")
	item2 := queueItem.NewQueueItem("error")
	item3 := queueItem.NewQueueItem("data3")
	item4 := queueItem.NewQueueItem("data4")
	item5 := queueItem.NewQueueItem("data5")

	q.Add(item1)
	q.Add(item2)
	q.Add(item3)
	q.Add(item4)
	q.Add(item5)
	fmt.Println(q)

	pool.Wait()
	fmt.Println(q)
}
