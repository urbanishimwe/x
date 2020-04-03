package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func readInput(t int) {
	var N, B int
	var line string
	fmt.Scanf("%d %d", &N, &B)
	A := make([]int, N)
	fmt.Scanf("%s", &line)
	lined := strings.Split(line, " ")
	fmt.Println(line, lined)
	for i, v := range lined {
		A[i], _ = strconv.Atoi(v)
	}
	logic(B, t, A)
}

func logic(B int, t int, A []int) {
	sort.Ints(A)
	used := 0
	allocated := 0
	for _, a := range A {
		used += a
		if used > B {
			fmt.Printf("Case #%d: %d\n", t, allocated)
			return
		}
		allocated++
	}
	fmt.Printf("Case #%d: %d\n", t, allocated)
}

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		readInput(i + 1)
	}
}
