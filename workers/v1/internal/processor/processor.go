package processor

import (
	"batch-v1/internal/config"
	"batch-v1/internal/queueItem"
	"fmt"
	"time"
)

type Processor struct {
	description string
}

func NewProcessor(description string) *Processor {
	return &Processor{
		description: description,
	}
}

func (p *Processor) ProcessItem(item *queueItem.QueueItem) error {
	// Process the item
	fmt.Printf("Processing item: %s\n", item.Data)
	// Simulate processing logic
	time.Sleep(config.SimulateProcessing)

	// Simulate error
	if item.Data == "error" {
		return fmt.Errorf("failed to process item: %s", item.Data)
	}
	item.Status = queueItem.ItemProcessed
	return nil
}

/*
Struct Definition:
Processor is defined as a struct with a single field description of type string.
Structs in Go are similar to objects in TypeScript, used to group related data together.
Constructor Function:
NewProcessor is a constructor function that returns a pointer to a new Processor instance.
The return type is *Processor, indicating a pointer to a Processor struct.
Using a pointer here allows the function to return a reference to the newly created Processor instance, which can be used to modify the instance directly.
Method with Pointer Receiver:
ProcessItem is a method with a pointer receiver *Processor.
The receiver p *Processor means that the method operates on the pointer to the Processor instance.
Using a pointer receiver allows the method to modify the Processor instance's fields if needed and ensures that the method operates on the same instance.

Pointers vs. Non-Pointers
Pointers:
Used when you need to modify the original variable or struct.
Avoids copying large data structures, improving performance.
Allows sharing of data between functions and methods.
Non-Pointers:
Used when you only need to read the value and do not need to modify the original variable.
Simpler and safer when you do not need to manage memory or modify the original data.
In JavaScript, you typically don't need to worry about pointers explicitly.
JavaScript handles memory management and references for you. When you pass objects or arrays to functions,
they are passed by reference, meaning that modifications to the object or array within the function will
affect the original object or array. Primitive types (like numbers and strings), however, are passed by value,
meaning that modifications within the function do not affect the original value.

In contrast, Go requires you to explicitly use pointers when you want to pass references to variables, allowing
you to modify the original data. This provides more control over memory management and can lead to more efficient
programs, but it also requires a deeper understanding of how memory and references work.

Pointer to Processor:
NewProcessor returns a pointer to a Processor instance, allowing the caller to modify the instance if needed.
ProcessItem method uses a pointer receiver to operate on the Processor instance directly.

In Go, when a function starts with parentheses containing a receiver (e.g., func (p *Processor)), it indicates that
the function is a method associated with a struct. The receiver specifies the type to which the method belongs.
In this example, ProcessItem is a method of the Processor struct.
*/
