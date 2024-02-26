package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(math.Floor(float64(int(temp))/10)) * 10
		groups[group] = append(groups[group], temp)
	}

	for group, temps := range groups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
