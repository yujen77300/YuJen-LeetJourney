package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1: nums = [2,7,11,15], target = 9",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2: nums = [3,2,4], target = 6",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3: nums = [3,3], target = 6",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			// only check if two arrays contain the same elements
			if !reflect.DeepEqual(result, tt.expected) && !reflect.DeepEqual([]int{result[1], result[0]}, tt.expected) {
				t.Errorf("twoSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}
