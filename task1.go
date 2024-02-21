package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) PrintAge() {
	fmt.Println(h.Age)
}

func (h *Human) PrintName() {
	fmt.Println(h.Name)
}

type Action struct {
	Human
	Some string
}

func (a *Action) PrintSome() {
	fmt.Println(a.Some)
}
func main() {
	human := Human{"Alex", 24}
	action := Action{human, "Some"}

	action.PrintAge()
	action.PrintName()

	action.PrintSome()
}
