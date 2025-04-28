package longestconsecutivesequence

import "testing"


func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
	}

	for _, tt := range tests {
		result := longestConsecutive(tt.nums)
		if result != tt.output {
			t.Errorf("For input %v expected %d but got %d", tt.nums, tt.output, result)
		}
	}
}