package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		} else {
			arr[right], arr[i] = arr[i], arr[right]
			right--
		}
	}
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	arr := []int{2, 5, 1, 3, 7, 213, 11, 42, 67, 9, 8}
	quickSort(arr)
	fmt.Println(arr)
}
