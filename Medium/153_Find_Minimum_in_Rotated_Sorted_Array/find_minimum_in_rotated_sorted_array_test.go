package findminimuminrotatedsortedarray

import "testing"

// Example 1:

// Input: nums = [3,4,5,1,2]
// Output: 1
// Explanation: The original array was [1,2,3,4,5] rotated 3 times.
// Example 2:

// Input: nums = [4,5,6,7,0,1,2]
// Output: 0
// Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
// Example 3:

// Input: nums = [11,13,15,17]
// Output: 11
// Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

func TestFinMin(t *testing.T) {
	test := []struct {
		nums []int
		min  int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}

	for _, tt := range test {
		t.Run("TestFinMin", func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.min {
				t.Errorf("findMin() = %v, want %v", got, tt.min)
			}
		})
	}
}
