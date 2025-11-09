package binarysearch

import "testing"

var testCases = []struct {
	input          []int
	valueToTest    int
	expectedResult bool
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 15}, 5, true},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 15}, 5, true},
	{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 14, true},
	{[]int{3, 6, 9, 12, 15, 18, 21}, 5, false},
	{[]int{0, 1, 1, 2, 3, 5, 8, 13, 21}, 8, true},
	{[]int{-10, -5, 0, 5, 10, 15, 20}, -5, true},
	{[]int{-3, -2, -1, 0, 1, 2, 3}, 4, false},
	{[]int{10, 20, 30, 40, 50, 60, 70}, 10, true},
	{[]int{1, 3, 5, 7, 9, 11, 13, 15}, 2, false},
	{[]int{100, 200, 300, 400, 500, 600, 700}, 700, true},
	{[]int{5, 15, 25, 35, 45, 55}, 60, false},
	{[]int{1, 2, 4, 8, 16, 32, 64}, 32, true},
	{[]int{11, 22, 33, 44, 55, 66, 77}, 44, true},
	{[]int{1, 3, 5, 7, 9, 11, 13}, 14, false},
	{[]int{-50, -25, 0, 25, 50, 75, 100}, 0, true},
	{[]int{9, 18, 27, 36, 45, 54, 63, 72}, 19, false},
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range testCases {
		res := BinarySearch(tt.input, tt.valueToTest)
		if res != tt.expectedResult {
			t.Errorf("got %v, want %v", res, tt.expectedResult)
			continue
		}
	}
}
