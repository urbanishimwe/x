// Problem description look in the PDF current directory, similar-strings.pdf
// still some improvement
package main

import (
	"fmt"
)

func condition(a string, b string) bool {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			c1 := a[i] == a[j] && b[i] == b[j]
			c2 := a[i] != a[j] && b[i] != b[j]
			if !(c1 || c2) {
				return false
			}
		}
	}
	return true
}

/*
 * Complete the similarStrings function below.
 */
func similarStrings(s string, n int, q int, queries [][2]int) (answer []int) {
	for i := 0; i < q; i++ {
		rl := s[queries[i][0]-1 : queries[i][1]]
		l := len(rl)
		if l == 1 {
			answer = append(answer, n)
			continue
		}
		var xy string
		var j, found int
		for l+j <= n {
			xy = s[j : l+j]
			if condition(rl, xy) {
				found++
			}
			j++
		}
		answer = append(answer, found)
	}
	return answer
}

func main() {
	var N, Q int
	var S string
	fmt.Scanf("%d %d", &N, &Q)
	fmt.Scan(&S)
	if len(S) != N {
		panic("bad input!")
	}
	Queries := make([][2]int, Q)
	for i := 0; i < Q; i++ {
		var r, l int
		fmt.Scanf("%d %d", &r, &l)
		Queries[i] = [2]int{r, l}
	}
	result := similarStrings(S, N, Q, Queries)

	for _, v := range result {
		fmt.Printf("%d\n", v)
	}
}
