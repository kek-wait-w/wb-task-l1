package main

import (
	"fmt"
)

func setBit(value int64, i int, bitValue int) int64 {
	mask := int64(1) << i

	if bitValue == 1 {
		return value | mask
	} else {
		return value &^ mask
	}
}

func main() {
	var value int64 = 56
	i := 1

	value = setBit(value, i, 1)
	fmt.Println(value)

	value = setBit(value, i, 0)
	fmt.Println(value)
}
