package main

import (
	"fmt"
	"time"
)

func main() {

	inputChan := make(chan int)
	outputChan := make(chan int)

	go func() {
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			inputChan <- num
		}
		close(inputChan)
	}()

	go func() {
		for num := range inputChan {
			outputChan <- num * num
			time.Sleep(time.Second)
		}
		close(outputChan)
	}()

	for result := range outputChan {
		fmt.Println(result)
	}
}
