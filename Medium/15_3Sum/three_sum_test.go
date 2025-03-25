package threesum

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if len(got) != len(tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
			for i, arr := range got {
				if len(arr) != len(tt.want[i]) {
					t.Fatalf("got: %v, want: %v", got, tt.want)
				}

				for j, val := range arr {
					if val != tt.want[i][j] {
						t.Fatalf("got: %v, want: %v", got, tt.want)
					}
				}
			}
		})
	}
}
