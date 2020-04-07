// Problem description look in the PDF current directory, similar-strings.pdf
package main

import (
	"bufio"
	"fmt"
	"os"
)

func condition(a string, b string) bool {
	l := len(a)
	j := l - 1
	for i := 0; i < j; i++ {
		c1 := a[i] == a[j] && b[i] == b[j]
		c2 := a[i] != a[j] && b[i] != b[j]
		if !(c1 || c2) {
			return false
		}
		j--
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
	file, _ := os.Open("input.in")
	stdin := bufio.NewReader(file)
	fmt.Fscanf(stdin, "%d %d", &N, &Q)
	fmt.Fscan(stdin, &S)
	if len(S) != N {
		panic("bad input!")
	}
	Queries := make([][2]int, Q)
	for i := 0; i < Q; i++ {
		var r, l int
		fmt.Fscanf(stdin, "\n%d %d", &r, &l)
		Queries[i] = [2]int{r, l}
	}
	result := similarStrings(S, N, Q, Queries)
	for _, v := range result {
		fmt.Printf("%d\n", v)
	}
}
