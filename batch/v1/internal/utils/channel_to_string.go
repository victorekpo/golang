package utils

import (
	"batch-v1/internal/workers/queue_item"
	"strings"
	"time"
)

type QueueItem = queue_item.QueueItem

func ChannelToString(ch chan *queue_item.QueueItem) string {
	var items []string
	timeout := time.After(100 * time.Millisecond)

	for {
		select {
		case item := <-ch:
			if item != nil {
				items = append(items, item.String())
			}
		case <-timeout:
			return strings.Join(items, ", ")
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func ConvertListToString(items []*QueueItem) string {
	var itemStrings []string
	for _, item := range items {
		itemStrings = append(itemStrings, item.String())
	}
	return strings.Join(itemStrings, ", ")
}

/*
Sending to a Channel: Use the <- operator to send a value to a channel.
ch <- value  // Sends 'value' to the channel 'ch'
Receiving from a Channel: Use the <- operator to receive a value from a channel.
value := <-ch  // Receives a value from the channel 'ch' and assigns it to 'value'

package main

import (
  "fmt"
)

func main() {
  ch := make(chan int)

  // Sending to the channel
  go func() {
    ch <- 42
  }()

  // Receiving from the channel
  value := <-ch
  fmt.Println(value)  // Output: 42
}

In this example:
The goroutine sends the value 42 to the channel ch.
The main function receives the value from the channel and prints it.
Yes, Go automatically removes the item from the channel once it is received by a worker.

In Go, channels can be in different states based on their usage and lifecycle. Here are the primary states of a channel:
Open: A channel is open when it is created using the make function and can send and receive values.
Closed: A channel is closed when the close function is called on it. After a channel is closed, no more values can be sent to it, but it can still be read until it is empty.
Empty: A channel is empty when there are no values to be received from it. This can happen either because no values have been sent to it yet, or all sent values have been received.
Full: A channel is full when its buffer is full, meaning it cannot accept any more values until some are received from it.
Here is a brief explanation of how these states interact:
Open and Empty: The channel is open but has no values to be received.
Open and Full: The channel is open and its buffer is full, so it cannot accept more values until some are received.
Closed and Empty: The channel is closed and all values have been received, so it cannot be used anymore.
Closed and Non-Empty: The channel is closed but still has values that can be received.

package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 2) // Create a buffered channel with capacity 2

    // Open and Empty
    fmt.Println("Open and Empty")
    select {
    case val := <-ch:
        fmt.Println("Received:", val)
    default:
        fmt.Println("Channel is empty")
    }

    // Open and Full
    ch <- 1
    ch <- 2
    fmt.Println("Open and Full")
    select {
    case ch <- 3:
        fmt.Println("Sent 3")
    default:
        fmt.Println("Channel is full")
    }

    // Closed and Non-Empty
    close(ch)
    fmt.Println("Closed and Non-Empty")
    for val := range ch {
        fmt.Println("Received:", val)
    }

    // Closed and Empty
    fmt.Println("Closed and Empty")
    select {
    case val := <-ch:
        fmt.Println("Received:", val)
    default:
        fmt.Println("Channel is empty and closed")
    }
}

Resolved Issue: The issue was likely caused by blocking reads from the channels while they were being written to, which could
lead to a deadlock. The updated ChannelToString function uses a non-blocking read approach with a timeout to avoid this.

Here's how it works:
Non-blocking Read with Timeout: The select statement is used to attempt to read from the channel. If an item is available,
it is processed and added to the items slice.
Timeout Case: If no item is available within 100 milliseconds, the time.After case is triggered, and the function returns
the concatenated string of items. This ensures that the function does not block indefinitely.
This approach prevents the function from getting stuck waiting for items from the channel, thus avoiding potential deadlocks.
*/
