package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	volume() float64
}

type Pole struct {
	radius, height float64
}

type Square struct {
	lng1, lng2, lng3 float64
}

func (p Pole) area() float64 {
	return math.Pi*p.radius*p.radius*2 + 2*math.Pi*p.radius*p.height
}

func (s Square) area() float64 {
	return 2 * (s.lng1*s.lng2 + s.lng1*s.lng3 + s.lng2*s.lng3)
}

func (p Pole) volume() float64 {
	return (math.Pi * p.radius * p.radius) * p.height
}

func (s Square) volume() float64 {
	return s.lng1 * s.lng2 * s.lng3
}

func main() {

	cy1 := Pole{10, 10}
	cy2 := Pole{4.2, 15.6}
	cu1 := Square{10.5, 20.2, 20}
	cu2 := Square{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(m ...geometry) {
	for _, val := range m {
		fmt.Printf("%.2f, %.2f\n", val.area(), val.volume())
	}
}
