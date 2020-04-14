package algorithm

// BinarySearch perform binary search, returns index of the element and bool which indicate
// if the element was found
func BinarySearch(array []int, key int) (index int, found bool) {
	if len(array) == 0 {
		return
	}
	var (
		mid   = len(array) / 2
		low   = 0
		high  = len(array)
	)
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
	return
}

// BubbleSort sort array in ascending order
func BubbleSort(array []int) []int {
	length := len(array)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < length; i++ {
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
				swapped = true
			}
		}
	}
	return array
}
