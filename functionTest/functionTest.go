package functionTest

import "fmt"

func BubbleSort(nums []int) []int {
	var temp int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	return nums
}

func InputNums() (result []int) {
	var count, num int
	fmt.Scanf("%d", &count)

	i := 0
	for i < count {
		fmt.Scanf("%d", &num)
		result = append(result, num)
		i++
	}

	return
}

func OutputNums(nums []int) {
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
}
