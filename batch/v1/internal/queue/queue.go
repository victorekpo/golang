package queue

import (
	"batch-v1/internal/queue_item"
	"batch-v1/internal/utils"
	"fmt"
	"strings"
	"sync"
)

type QueueItem = queue_item.QueueItem

type Queue struct {
	Items []*QueueItem
	Queue chan *QueueItem
	DLQ   chan *QueueItem
	Wg    *sync.WaitGroup
}

func NewQueue(bufferSize int, wg *sync.WaitGroup) *Queue {
	return &Queue{
		Items: []*QueueItem{},
		Queue: make(chan *QueueItem, bufferSize),
		DLQ:   make(chan *QueueItem, bufferSize),
		Wg:    wg,
	}
}

func (q *Queue) Add(item *QueueItem) {
	q.Items = append(q.Items, item)
	q.Wg.Add(1)
	q.Queue <- item
}

func (q *Queue) Remove(item *QueueItem) {
	for i, work := range q.Items {
		if work.ID == item.ID {
			q.Items = append(q.Items[:i], q.Items[i+1:]...)
			break
		}
	}
}

func (q *Queue) Length() int {
	return len(q.Items)
}

func (q *Queue) String() string {
	var items []string
	for _, item := range q.Items {
		items = append(items, item.String())
	}

	queueItems := utils.ChannelToString(q.Queue)
	dlqItems := utils.ChannelToString(q.DLQ)

	return fmt.Sprintf("Queue length: %d, Items: [%s], Queue: [%s], DLQ: [%s]", q.Length(), strings.Join(items, ", "), queueItems, dlqItems)
}

func (q *Queue) QueueItems() string {
	return fmt.Sprintf("Queue: [%s]", utils.ChannelToString(q.Queue))
}

func (q *Queue) DLQItems() string {
	return fmt.Sprintf("DLQ: [%s]", utils.ChannelToString(q.DLQ))
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
