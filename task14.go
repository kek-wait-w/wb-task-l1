package main

import "fmt"

func determinate(expr interface{}) {
	if _, ok := expr.(string); ok {
		fmt.Println("String type")
	}
	if _, ok := expr.(int); ok {
		fmt.Println("Int type")
	}
	if _, ok := expr.(bool); ok {
		fmt.Println("Bool type")
	}
	if _, ok := expr.(chan (int)); ok {
		fmt.Println("Chan type")
	}
}
func main() {
	val1 := "Str"
	determinate(val1)
	val2 := 1
	determinate(val2)
	val3 := true
	determinate(val3)
	val4 := make(chan int)
	determinate(val4)
}
