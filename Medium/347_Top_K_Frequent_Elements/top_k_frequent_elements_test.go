package topkfrequentelements

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	test := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{4, 2, 2, 2, 1, 1, 1}, 2, []int{1, 2}},
	}

	for _, tt := range test {
		result := topKFrequent(tt.nums, tt.k)
		if len(result) != len(tt.expected) {
			t.Errorf("Expected length %d, but got %d", len(tt.expected), len(result))
			continue
		}

		// Check if each expected element is in the result
		for _, expected := range tt.expected {
			if !slices.Contains(result, expected) {
				t.Errorf("Expected %v to contain %d, but it does not", result, expected)
			}
		}
	}
}
