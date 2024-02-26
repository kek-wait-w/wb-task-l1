package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало программы")
	sleep(3)
	fmt.Println("Прошло 3 секунды")
}
