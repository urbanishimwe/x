package hackerrank

import (
	"fmt"
)

// question description at https://www.hackerrank.com/challenges/apple-and-orange/problem
func apple() {
	var S, T, A, B, M, N int
	fmt.Scan(&S, &T, &A, &B, &M, &N)
	Apple := make([]int, M)
	Orange := make([]int, N)

	appleOnHouse, orangeOnHouse := 0, 0
	for i := range Apple {
		fmt.Scan(&Apple[i])
		if (Apple[i]+A) >= S && (Apple[i]+A) <= T {
			appleOnHouse++
		}
	}
	for i := range Orange {
		fmt.Scan(&Orange[i])
		if (Orange[i]+B) <= T && (Orange[i]+B) >= S {
			orangeOnHouse++
		}
	}
	fmt.Println(appleOnHouse)
	fmt.Println(orangeOnHouse)
}
