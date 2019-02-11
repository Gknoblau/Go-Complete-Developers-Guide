package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	x float64
	y float64
}
type square struct {
	x float64
}

func main() {
	s := square{x: 10}
	printArea(s)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s triangle) getArea() float64 {
	return s.x * s.y * 1 / 2
}

func (s square) getArea() float64 {
	return s.x * s.x
}
