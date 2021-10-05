package breakContinue

import "fmt"

func BreakContinue() {
	for i := 2; i <= 9; i++ {
		if i%2 == 0 {
			continue
		} else {
			for j := 1; j <= 9; j++ {
				fmt.Printf("%d x %d = %d\n", i, j, i*j)
				if j == i {
					break
				}
			}
			fmt.Println()
		}
	}
}
