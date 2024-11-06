package queue_item

import (
	"fmt"
	"github.com/google/uuid"
)

type ItemStatus int

const (
	ItemPending ItemStatus = iota
	ItemProcessing
	ItemProcessed
)

func (s ItemStatus) String() string {
	return [...]string{"pending", "processing", "processed"}[s]
}

type QueueItem struct {
	ID     string
	Data   interface{}
	Status ItemStatus
}

func NewQueueItem(data interface{}) *QueueItem {
	return &QueueItem{
		ID:     uuid.New().String(),
		Data:   data,
		Status: ItemPending,
	}
}

func (item *QueueItem) String() string {
	return fmt.Sprintf("Data: %s Status: %s", item.Data, item.Status)
}
