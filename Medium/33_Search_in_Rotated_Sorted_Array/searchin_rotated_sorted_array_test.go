package searchinrotatedsortedarray

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
			in:       []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Example 2",
			in:       []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "Example 3",
			in:       []int{1},
			target:   0,
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
