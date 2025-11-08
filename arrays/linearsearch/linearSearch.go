package linearsearch

func linearSearch(arr []int, val int) bool {
	for _, value := range arr {
		if val == value {
			return true
		}
	}
	return false
}
