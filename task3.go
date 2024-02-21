package main

import (
	"fmt"
	"sync"
)

func calc(num int, wg *sync.WaitGroup, res chan int) {
	defer wg.Done()
	square := num * num
	res <- square
}

func main() {
	buf := 0
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(len(numbers))

	for _, num := range numbers {
		go calc(num, wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		buf += result
	}
	fmt.Println(buf)
	wg.Wait()
}
