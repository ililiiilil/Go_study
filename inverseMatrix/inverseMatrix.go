package inverseMatrix

import "fmt"

func InverseMatrix() {
	var A = [2][2]int{
		{7, 3},
		{5, 2},
	}

	d := A[0][0]*A[1][1] - A[1][0]*A[0][1]

	if d != 0 {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				A[i][j] = A[i][j] / d
			}
		}

		fmt.Println("true")
		fmt.Printf("%d %d\n", A[1][1], -1*A[0][1])
		fmt.Printf("%d %d\n", -1*A[1][0], A[0][0])
	} else {
		fmt.Println("false")
	}
}
