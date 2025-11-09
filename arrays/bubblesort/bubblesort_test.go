package bubblesort

import "testing"

var testCases = []struct {
	input          []int
	expectedResult []int
}{
	{[]int{1, 3, 2, 5, 2, 4}, []int{1, 2, 2, 3, 4, 5}},
	{[]int{1, 3, 2, 5, 2, 4}, []int{1, 2, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{10, 9, 8, 7, 6, 5}, []int{5, 6, 7, 8, 9, 10}},
	{[]int{2, 1}, []int{1, 2}},
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{4, 5, 3, 1, 2}, []int{1, 2, 3, 4, 5}},
	{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{100, 50, 200, 150, 0}, []int{0, 50, 100, 150, 200}},
	{[]int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
	{[]int{1}, []int{1}},
	{[]int{2, 0, -1, 5, 3}, []int{-1, 0, 2, 3, 5}},
	{[]int{-5, -10, 0, 5, 10}, []int{-10, -5, 0, 5, 10}},
	{[]int{7, 3, 9, 1, 6}, []int{1, 3, 6, 7, 9}},
	{[]int{4, 4, 2, 2, 1, 1}, []int{1, 1, 2, 2, 4, 4}},
	{[]int{10, -2, 33, 5, 0, -7}, []int{-7, -2, 0, 5, 10, 33}},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range testCases {
		res := BubbleSort(tt.input)
		for in, val := range res {
			if val != tt.input[in] {
				t.Errorf("got %v, want %v", res, tt.expectedResult)
				continue
			}
		}
	}
}
