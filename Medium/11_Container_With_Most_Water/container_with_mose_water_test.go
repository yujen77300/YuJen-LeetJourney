package containerwithmostwater

import "testing"


func TestMaxArea(t *testing.T) {
	test := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "Example 1",
			in:       []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			in:       []int{1, 1},
			expected: 1,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.in)
			if result != tt.expected {
				t.Fatalf("got: %v, want: %v", result, tt.expected)
			}
		})
	}

}
