package main

import (
	"context"
	"fmt"
	"time"
)

func worker1(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			fmt.Println("Worker1 working")
			time.Sleep(time.Second)
		}
	}
}

func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Worker2 working")
			time.Sleep(time.Second)
		}
	}
}

func worker3(stop *bool) {
	for !*stop {
		fmt.Println("Worker3 working")
		time.Sleep(time.Second)
	}
}

func main() {
	stop1 := make(chan bool)
	go worker1(stop1)

	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)

	stop3 := false
	go worker3(&stop3)

	time.Sleep(3 * time.Second)
	stop1 <- true
	cancel()
	stop3 = true

	fmt.Println("Done")
}
