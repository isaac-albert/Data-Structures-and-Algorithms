package twocrystalballs

import "testing"

var testCases = []struct {
	input          []int
	expectedResult int
}{
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}, 9},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}, 9},
	{[]int{0, 0, 0, 0, 1, 1, 1, 1, 1}, 5},
	{[]int{0, 0, 1, 1, 1, 1, 1, 1}, 3},
	{[]int{0, 0, 0, 0, 0, 0, 1}, 7},
	{[]int{0, 1, 1, 1, 1, 1}, 2},
	{[]int{1, 1, 1, 1, 1, 1, 1}, 1},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 10},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 13},
	{[]int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1}, 4},
	{[]int{0, 0, 0, 0, 0, 1, 1, 1, 1}, 6},
	{[]int{0, 0, 0, 0, 1}, 5},
	{[]int{0, 0, 0, 0, 0, 0, 0, 1}, 8},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 11},
	{[]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}, 6},
	{[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1}, 8},
}

func TestTwoCrystallBalls(t *testing.T) {
	for _, tt := range testCases {
		res := TwoCrystalBalls(tt.input)
		if res != tt.expectedResult {
			t.Errorf("got %v, want %v", res, tt.expectedResult)
			continue
		}
	}
}
