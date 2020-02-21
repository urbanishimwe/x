package hackerrank

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
func dynamicArray(n int, queries [][]int) (answer []int) {
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
