package algorithm

// BinarySearch perform binary search on an array returns -1 if key is not found
func BinarySearch(array []int, key int) (index int) {
	if len(array) == 0 {
		return -1
	}
	var (
		mid   = len(array) / 2
		low   = 0
		high  = len(array)
		found = false
	)
	index = -1
	for !found {
		if array[mid] > key {
			high = mid
			mid = high / 2
		} else if array[mid] < key {
			low = mid + 1
			mid = (high + low) / 2
		} else if high == low {
			found = true
		} else {
			index = mid
			break
		}
	}
	return index
}
