package main

import (
	"fmt"
	"strings"
)

func uniq(str string) bool {
	uni := make(map[string]int)
	lowerStr := strings.ToLower(str)
	s := strings.Split(lowerStr, "")
	for _, i := range s {
		uni[i]++
		if uni[i] > 1 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(str1, ":", uniq(str1))
	fmt.Println(str2, ":", uniq(str2))
	fmt.Println(str3, ":", uniq(str3))
}
