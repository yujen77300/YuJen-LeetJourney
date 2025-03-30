package maximumproductsubarray

import "testing"

func TestMaxProduct(t *testing.T) {
	test := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "Example 1",
			in:       []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "Example 2",
			in:       []int{-2, 0, -1},
			expected: 0,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProduct(tt.in)
			if result != tt.expected {
				t.Fatalf("got: %v, want: %v", result, tt.expected)
			}
		})
	}
}
