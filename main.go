package main

import "fmt"

func main() {
	var test = make([]string, 0)
	var name string

	for {
		fmt.Scanln(&name)
		if name != "0" {
			test = append(test, name)
		} else {
			break
		}
	}

	for i := 0; i < len(test); i++ {
		defer fmt.Println(test[i])
	}
}
