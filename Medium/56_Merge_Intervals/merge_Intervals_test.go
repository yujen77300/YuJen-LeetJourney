package mergeintervals

import "testing"


func TestMerge(t *testing.T) {
	test := []struct {
		name string
		in   [][]int
		expected  [][]int
	}{
		{
			name: "Example 1",
			in:   [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "Example 2",
			in:   [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name: "Example 3",
			in:   [][]int{{1, 4}, {0, 1}},
			expected:  [][]int{{0, 4}},
		},
		{
			name: "Example 4",
			in:   [][]int{{1, 4}, {2, 3}},
			expected:  [][]int{{1, 4}},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.in)
			if len(result) != len(tt.expected) {
				t.Fatalf("got: %v, want: %v", result, tt.expected)
			}
			for i, arr := range result {
				if len(arr) != len(tt.expected[i]) {
					t.Fatalf("got: %v, want: %v", result, tt.expected)
				}

				for j, val := range arr {
					if val != tt.expected[i][j] {
						t.Fatalf("got: %v, want: %v", result, tt.expected)
					}
				}
			}
		})
	}

}
