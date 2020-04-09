// For Problem description look in the PDF similar-strings.pdf in the current directory
package main

import (
	"fmt"
	"strings"
)

var Answer []string
var Found int

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
func similarStrings(s string, n int, q int, queries [][2]int) {
	for i := 0; i < q; i++ {
		go func(i int) {
			rl := s[queries[i][0]-1 : queries[i][1]]
			l := len(rl)
			if l == 1 {
				Answer[i] = fmt.Sprintf("%d", n)
				Found++
				return
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
			Answer[i] = fmt.Sprintf("%d", found)
			Found++
		}(i)
	}
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
	Answer = make([]string, Q)
	for i := 0; i < Q; i++ {
		var r, l int
		fmt.Scanf("%d %d", &r, &l)
		Queries[i] = [2]int{r, l}
	}
	similarStrings(S, N, Q, Queries)
	for {
		if Found == Q {
			break
		}
	}
	fmt.Print(strings.Join(Answer, "\n"))
}
