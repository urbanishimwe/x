// problem statement https://www.hackerrank.com/challenges/extra-long-factorials/problem
package hackerrank

import (
	"fmt"
	"math/big"
)

func extra_long_factorials() {
	var N int
	fmt.Scan(&N)
	var ans = big.NewInt(1)
	for i := N; i > 1; i-- {
		ans.Mul(ans, big.NewInt(int64(i)))
	}
	fmt.Println(ans.String())
}
