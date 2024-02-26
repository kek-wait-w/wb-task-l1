package main

import "fmt"

func main() {

	str := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, s := range str {
		set[s] = true
	}

	for s := range set {
		fmt.Println(s)
	}
}
