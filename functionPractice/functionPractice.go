package functionPractice

import "fmt"

func Guide() {
	fmt.Println("안녕하세요")
}

func Input() (int, int) {
	var a, b int
	fmt.Scanln(&a, &b)
	return a, b
}

func Multi(a, b int) int {
	return a * b
}

func PrintResult(num int) {
	fmt.Printf("결과값은 %d입니다.\n", num)
}
