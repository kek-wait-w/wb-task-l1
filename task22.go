package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Pow(2, 21)
	b := math.Pow(2, 22)

	if a <= math.Pow(2, 20) || b <= math.Pow(2, 20) {
		fmt.Println("Число меньше или равно 2^20")
		return
	}

	multiplication := a * b

	division := a / b

	addition := a + b

	subtraction := a - b

	fmt.Println("Умножение: ", multiplication)
	fmt.Println("Деление: ", division)
	fmt.Println("Сложение: ", addition)
	fmt.Println("Вычитание: ", subtraction)
}
