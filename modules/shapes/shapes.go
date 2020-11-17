package main

import (
	"fmt"
	"math"
)

type Shape interface {
	GetArea() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Square struct {
	SideLength float64
}

func (t Triangle) GetArea() float64 {
	return 0.5 * t.Base * t.Height
}

func (s Square) GetArea() float64 {
	return math.Pow(s.SideLength, 2)
}

func PrintArea(s Shape) {
	fmt.Println(s.GetArea())
}

func init() {
	t := Triangle{
		Height: 1.5,
		Base:   3,
	}

	PrintArea(t)

	s := Square{
		SideLength: 5,
	}

	PrintArea(s)
}
