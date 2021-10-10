package main

import "fmt"

func add(num1 int, num2 int, c chan int) {
	c <- num1 + num2
}

func main() {
	var num1, num2 int
	ch := make(chan int)

	fmt.Scanln(&num1, &num2)

	go add(num1, num2, ch)

	fmt.Println(<-ch)
}
