package queueItem

import (
	"fmt"
	"github.com/google/uuid"
)

type ItemStatus string

const (
	ItemPending    ItemStatus = "pending"
	ItemProcessing ItemStatus = "processing"
	ItemProcessed  ItemStatus = "processed"
)

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
	return fmt.Sprintf("ID: %s Data: %s Status: %s", item.ID, item.Data, item.Status)
}
