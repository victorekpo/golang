package queue

import (
	"batch-v1/internal/utils"
	"batch-v1/internal/workers/queue_item"
	"fmt"
)

type Item = queue_item.QueueItem

type Queue struct {
	Items     []*Item
	QueueList []*Item
	DLQList   []*Item
	QueueChan chan *Item
	DLQChan   chan *Item
}

func NewQueue(bufferSize int) *Queue {
	return &Queue{
		Items:     []*Item{},
		QueueList: []*Item{},
		DLQList:   []*Item{},
		QueueChan: make(chan *Item, bufferSize),
	}
}

func (q *Queue) Add(item *Item) {
	// Add all items to the Items list (for logging purposes)
	q.Items = append(q.Items, item)
	// Add item to the Queue list (for logging purposes)
	q.QueueList = append(q.QueueList, item)
	fmt.Printf("*** Current # items in queue: %d ***\n", q.QueueLength())
	// Add item to the Queue channel (for processing by workers)
	q.QueueChan <- item
}

func (q *Queue) AddDLQ(item *Item) {
	// Add item to the DLQ list (for logging purposes)
	q.DLQList = append(q.DLQList, item)
	fmt.Printf("*** Current # items in DLQ: %d ***\n", q.DLQLength())
}

func (q *Queue) ProcessAllDLQ() {
	fmt.Println("*** Processing all items in DLQ ***")
	// Process all items in the DLQ
	for _, item := range q.DLQList {
		q.Add(item)
	}
	q.DLQList = []*Item{}
}

func (q *Queue) Remove(item *Item) {
	// Remove item from the Queue list
	for i, work := range q.QueueList {
		if work.ID == item.ID {
			q.QueueList = append(q.QueueList[:i], q.QueueList[i+1:]...)
			break
		}
	}
	fmt.Printf("*** Current # items in queue: %d ***\n", q.QueueLength())
}

func (q *Queue) Length() int {
	return len(q.Items)
}

func (q *Queue) QueueLength() int {
	return len(q.QueueList)
}

func (q *Queue) DLQLength() int {
	return len(q.DLQList)
}

func (q *Queue) String() string {
	return fmt.Sprintf("Items length: %d, Items: [%s], Queue: [%s], DLQ: [%s]",
		q.Length(),
		utils.ConvertListToString(q.Items),
		utils.ConvertListToString(q.QueueList),
		utils.ConvertListToString(q.DLQList))
}

func (q *Queue) QueueItems() string {
	return fmt.Sprintf("Queue: [%s]", utils.ConvertListToString(q.QueueList))
}

func (q *Queue) DLQItems() string {
	return fmt.Sprintf("DLQ: [%s]", utils.ConvertListToString(q.DLQList))
}

/*
Explanation of Buffer Size for Channels
The buffer size for channels in Go is important because it determines how many items can be stored in the channel
without blocking the sender. When the buffer is full, any additional sends will block until space is available.
Conversely, if the buffer is empty, any receives will block until an item is available.

In this case, setting the buffer size equal to the number of workers ensures that each worker can process an
item without blocking the main goroutine that is adding items to the channel. This helps maintain smooth
processing and avoids potential deadlocks.
*/
