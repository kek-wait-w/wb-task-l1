package main

import "fmt"

func main() {
	arr1 := []string{"cat", "dog", "car", "new", " write", "string"}
	arr2 := []string{"dog", "winter", "house", "string"}
	intersection := make(map[string]int)

	for _, item := range arr1 {
		intersection[item] = 0
	}

	for str := range intersection {
		for _, s := range arr2 {
			if s == str {
				intersection[str]++
			}
		}
	}

	fmt.Println("Intersection of two sets:")
	for str := range intersection {
		if intersection[str] > 0 {
			fmt.Print(str, " ")
		}
	}
}
