// Problem description look in the PDF current directory, similar-strings.pdf
// still some improvement
package hackerrank

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
