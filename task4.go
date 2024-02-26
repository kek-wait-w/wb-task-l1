package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func worker(num int, info <-chan int) {
	for inf := range info {
		fmt.Println("Worker ", num, " : ", inf)
	}
}

func main() {

	mainChan := make(chan int)
	numWorkers, _ := strconv.Atoi(os.Args[1])
	for i := 1; i <= numWorkers; i++ {
		go worker(i, mainChan)
	}

	go func() {
		jobCounter := rand.Int()
		for {
			mainChan <- jobCounter
			jobCounter = rand.Int()
			time.Sleep(time.Second)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	close(mainChan)
}
