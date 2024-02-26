package main

import "fmt"

func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = removeIndex(arr, 2)
	fmt.Println(arr)
}
