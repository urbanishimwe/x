package hackerrank

import (
	"strconv"
	"strings"
)

func match(sub string, v string) bool {
	var x int
	for i := 0; i < len(sub); i++ {
		if sub[i] != v[i] {
			x++
		}
		if x > 1 {
			return false
		}
	}
	return true
}

/*
 * Complete the virusIndices function below.
 */
func virusIndices(p string, v string) string {
	lenP, lenV := len(p), len(v)
	var s string
	if lenP < lenV {
		return "No Match!"
	}
	var i int
	for (i + lenV) <= lenP {
		sub := p[i:(lenV + i)]
		if match(sub, v) {
			s = s + " " + strconv.Itoa(i)
		}
		i++
	}
	if s == "" {
		return "No Match!"
	}
	return strings.Trim(s, " ")
}
