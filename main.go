package main

import "fmt"

func main() {
	var test = make(map[string]int)
	var subject string
	var score int
	var sum int

	for {
		fmt.Scanf("%d %d", &subject, &score)
		test[subject] = score

		if test["b"] == 10 {
			break
		}
	}

	for _, s := range test {
		sum += s
	}

	for i, j := range test {
		fmt.Printf("%s %d\n", i, j)
	}
	fmt.Println(sum / len(test))
}
