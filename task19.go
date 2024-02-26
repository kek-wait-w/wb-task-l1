package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	var result string
	s := strings.Split(str, "")
	left, right := 0, len(s)-1

	if (len(s) % 2) == 0 {
		for i := 0; i < (len(s))/2; i++ {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
	} else {
		for i := 0; i < (len(s)-1)/2; i++ {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
	}

	for _, i := range s {
		result += i
	}
	return result
}

func main() {
	str := "абырвалг"
	fmt.Println(reverse(str))
}
