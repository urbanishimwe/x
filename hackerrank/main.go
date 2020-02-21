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
