package closer

import "fmt"

func calCoin(count, coin int) int {
	return count * coin
}

func Closer() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)

	if coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0 {
		fmt.Println("잘못된입력입니다.")
	} else {
		add10 := calCoin(coin10, 10)
		add50 := calCoin(coin50, 50)
		add100 := calCoin(coin100, 100)
		add500 := calCoin(coin500, 500)

		fmt.Print(add10 + add50 + add100 + add500)
	}
}
