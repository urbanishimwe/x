package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var S []string
var wg sync.WaitGroup

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

func main() {
	var T int
	fmt.Scan(&T)
	wg.Add(T)
	S = make([]string, T)
	for i := 0; i < T; i++ {
		var P, V string
		fmt.Scan(&P, &V)
		go func(i int) {
			s := virusIndices(P, V)
			S[i] = s
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(strings.Join(S, "\n"))
}
