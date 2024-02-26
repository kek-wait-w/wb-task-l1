package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(q *Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	distance := p1.Distance(p2)

	fmt.Println(distance)
}
