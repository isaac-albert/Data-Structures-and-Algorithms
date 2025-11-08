package linearsearch

import "testing"

var testCases = []struct {
	input          []int
	valueToTest    int
	expectedOutput bool
}{
	{[]int{1, 2, 3, 4, 5}, 5, true},
	{[]int{24, 56, 78, 98, 60, 59, 68, 56, 5, 6, 3, 9, 10}, 100, false},
	{[]int{4, 7, 1, 9, 2}, 100, false},
	{[]int{4, 7, 1, 9, 2}, 9, true},
	{[]int{10, 20, 30, 40, 50}, 35, false},
	{[]int{99, 45, 32, 10, 8}, 99, true},
	{[]int{1, 2, 3, 4, 5}, 5, true},
	{[]int{7, 7, 7, 7, 7}, 7, true},
	{[]int{-5, -10, 0, 5, 10}, -10, true},
	{[]int{-1, 2, -3, 4, -5, 6}, 4, true},
	{[]int{}, 10, false},
	{[]int{42}, 42, true},
	{[]int{99}, 100, false},
	{[]int{5, 10, 15, 20, 25, 30}, 15, true},
	{[]int{100, 90, 80, 70, 60, 50}, 75, false},
	{[]int{1000, 5000, 10000, 15000}, 5000, true},
	{[]int{24, 56, 78, 98, 60, 59, 68, 56, 5, 6, 3, 9, 10}, 68, true},
	{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 2, true},
}

func TestLinearSearc(t *testing.T) {
	for _, tt := range testCases {
		res := linearSearch(tt.input, tt.valueToTest)
		if res != tt.expectedOutput {
			t.Errorf("got %v, want %v", res, tt.expectedOutput)
		}
	}
}
