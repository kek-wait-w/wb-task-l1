package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex
	data := make(map[string]int)

	numWorkers := 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			key := fmt.Sprintf("key%d", workerID)
			data[key] = workerID
			fmt.Printf("Worker %d wrote data to map: %s -> %d\n", workerID, key, workerID)
		}(i)
	}

	wg.Wait()

	fmt.Println(data)
}
