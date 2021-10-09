package main

import "fmt"

var g float32 = 9.8

type object struct {
	m, v, h, ke, ep float32
}

func (o *object) Ep() float32 {
	return o.m * o.v * o.v / 2
}

func (o *object) Ke() float32 {
	return o.m * g * o.h
}

func main() {
	var n int
	var m, v, h float32

	objects := []object{}

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&m, &v, &h)
		o := object{m, v, h, 0, 0}
		o.ke = o.Ke()
		o.ep = o.Ep()
		objects = append(objects, o)
	}

	for _, o := range objects {
		fmt.Println(o.ep, o.ke, o.ke+o.ep)
	}
}
