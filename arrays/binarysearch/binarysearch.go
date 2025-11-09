package binarysearch

func BinarySearch(arr []int, value int) bool {
	low := 0
	high := len(arr)

	return search(arr, value, low, high)
}

func search(arr []int, value, low, high int) bool {

	for low < high {
		med := int(low + (high-low)/2)

		val := arr[med]
		if val == value {
			return true
		} else if value > val {
			low = med + 1
		} else {
			high = med
		}
	}

	return false
}
