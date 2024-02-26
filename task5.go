package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func sendler(ch chan int) {
	for {
		info := rand.Int()
		ch <- info
		time.Sleep(1 * time.Second)
	}
}

func resiver(ch chan int, done chan bool) {
	for {
		select {
		default:
			fmt.Println(<-ch)
		case <-done:
			return
		}
	}
}

func main() {

	arg, _ := strconv.Atoi(os.Args[1])
	workTime := time.Duration(arg)

	mainChan := make(chan int)
	done := make(chan bool)

	go sendler(mainChan)

	go resiver(mainChan, done)

	time.Sleep(workTime * time.Second)
	done <- true
}
