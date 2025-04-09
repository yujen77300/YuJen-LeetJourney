package numberofislands

import "testing"

func TestNumIslands(t *testing.T) {
	test := []struct {
		name     string
		in       [][]byte
		expected int
	}{
		{
			name: "Example 1",
			in: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			in: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := numIslands(tt.in)
			if result != tt.expected {
				t.Fatalf("got: %v, want: %v", result, tt.expected)
			}
		})
	}
}
