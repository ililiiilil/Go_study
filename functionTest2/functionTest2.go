package funcTiontest2

import "fmt"

func InputNums() []int {
	var score int
	var nums []int

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &score)
		nums = append(nums, score)
	}

	return nums
}

func CalExam(arr ...int) (ans, count1, count2 int) {
	ans, count1, count2 = 0, 0, 0
	for _, val := range arr {
		ans += val
		if val >= 90 {
			count1 += 1
		} else if val < 70 {
			count2 += 1
		}
	}
	return
}

func PrintResult(ans, count1, count2 int) {
	var result bool = true

	if ans < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	}
	if count1 < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	}
	if count2 > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}

	if !result {
		fmt.Println("아이패드를 살 수 없습니다.")
	} else {
		fmt.Println("아이패드를 살 수 있습니다.")
	}
}
