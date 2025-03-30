package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	test := []struct {
		name     string
		in       []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			in:       []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "Example 2",
			in:       []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.in, tt.target)
			if result != tt.expected {
				t.Fatalf("got: %v, want: %v", result, tt.expected)
			}
		})
	}
}
