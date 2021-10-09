package main

import "fmt"

func FunctionTest3() {

	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func(m, v float32) float32 {
		return m * v * v / 2
	}
	potEnergy := func(m, h float32) float32 {
		return m * h * 9.8
	}

	ke := kinEnergy(m, v)
	pe := potEnergy(m, h)
	fmt.Println(ke, pe, ke+pe)
}
