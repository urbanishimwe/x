package hackerrank

import (
	"strconv"
	"strings"
)

// ReverseArray in a very optim way
func ReverseArray(a []int) []int {
	length := len(a)
	i := 0
	j := length - 1
	mid := length / 2
	for i < mid && j >= mid {
		temp := a[i]
		a[i] = a[j]
		a[j] = temp
		j--
		i++
	}
	return a
}

//DynamicArray perform queries on dynamic array
func DynamicArray(n int, queries [][]int) (answer []int) {
	lastAnswer := int(0)
	seqList := make([][]int, n)
	for _, value := range queries {
		N, x, y := value[0], value[1], value[2]
		seq := (x ^ lastAnswer) % n
		if N == 1 {
			seqList[seq] = append(seqList[seq], y)
		} else {
			lastAnswer = seqList[seq][y%len(seqList[seq])]
			answer = append(answer, lastAnswer)
		}
	}
	return answer
}

// LeftRotation A left rotation operation on an array of size n shifts each of the array's elements 1 unit to the left. For example, if 2(d) left rotations are performed on array [1, 2, 3, 4, 5], then the array would become [3, 4, 5, 1, 2].
func LeftRotation(array []int, n int, d int) []int {
	// The complexity here is d*(n - 1)
	for i := 0; i < d; i++ {
		swap := array[0]
		for j := 0; j < n-1; j++ {
			array[j] = array[j+1]
		}
		array[n-1] = swap
	}
	return array
}

// SparseArrays strings is array to query from, queries is array of those queries to search from strings
func SparseArrays(strings []string, queries []string) (answer []int) {
	for _, query := range queries {
		found := 0
		for _, matchWith := range strings {
			if query == matchWith {
				found++
			}
		}
		answer = append(answer, found)
	}
	return answer
}

// ArrayManipulation Array Manipulation challenges for small dataset
func ArrayManipulation(n int, queries [][]int) (answer int) {
	lastArray := make([]int, n)
	for _, value := range queries {
		a, b, k := value[0], value[1], value[2]
		for i := a - 1; i < b; i++ {
			lastArray[i] += k
			if lastArray[i] > answer {
				answer = lastArray[i]
			}
		}
	}
	return answer
}

// TimeConversion Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.
func TimeConversion(s string) string {
	s = strings.Trim(s, " \n\r")
	amp := s[len(s)-2:]
	h := s[0:2]
	s = strings.Replace(s, amp, "", 1)
	if amp == "AM" {
		if h == "12" {
			return strings.Replace(s, h, "00", 1)
		}
		return s
	}
	if h == "12" {
		return s
	}
	i, _ := strconv.Atoi(h)
	return strings.Replace(s, h, strconv.Itoa(i+12), 1)
}

// GradingStudents https://www.hackerrank.com/challenges/grading/problem
func GradingStudents(gr []int) []int {
	r := make([]int, len(gr))
	for _, v := range gr {
		if v >= 35 {
			mod := v % 5
			nx := v + 5 - mod
			if (nx - v) < 3 {
				v = nx
			}
		}
		r = append(r, v)
	}
	return r
}
