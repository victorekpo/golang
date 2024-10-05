package main

import (
  "fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %v: Finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	var maxWorkers = 5

	for i := 0; i < maxWorkers; i++ {
		fmt.Println("Adding worker", i)
		wg.Add(1)
		go worker(&wg, i)
	}

	fmt.Printf("Waiting for %d workers to finish\n", maxWorkers)
	wg.Wait()
	fmt.Println("All Workers Completed")
}
