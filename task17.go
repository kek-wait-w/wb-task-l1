package main

import "fmt"

func binarySearch(arr []int, key int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		midVal := arr[mid]
		if midVal == key {
			return mid
		}
		if midVal > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 5, 7, 8, 9, 11, 42, 67, 213}

	fmt.Println(binarySearch(arr, 5))

}
