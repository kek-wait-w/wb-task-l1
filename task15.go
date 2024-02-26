package main

import (
	"sync"
)

var (
	//Так как переменная глобальная возможен доступ к ней из разных горутин в одно и тоже время
	justString string
	//Добавляем Mutex для корректной работы
	mutex sync.Mutex
)

func someFunc() {
	v := createHugeString(1 << 10)
	mutex.Lock()
	justString = v[:100]
	mutex.Unlock()
}

func main() {
	someFunc()
}
