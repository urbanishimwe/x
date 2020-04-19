// problem statement https://www.hackerrank.com/challenges/extra-long-factorials/problem
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N int
	fmt.Scan(&N)
	var ans = big.NewInt(1)
	for i := N; i > 1; i-- {
		ans = ans.Mul(ans, big.NewInt(int64(i)))
	}
	fmt.Println(ans.Text(10))
}
