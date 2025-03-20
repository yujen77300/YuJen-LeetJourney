package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	test := []struct {
		name      string
		prices    []int
		maxProfit int
	}{
		{
			name:      "Example 1",
			prices:    []int{7, 1, 5, 3, 6, 4},
			maxProfit: 5,
		},
		{
			name:      "Example 2",
			prices:    []int{7, 6, 4, 3, 1},
			maxProfit: 0,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.prices)
			if result != tt.maxProfit {
				t.Errorf("maxProfit() = %v, want %v", result, tt.maxProfit)
			}
		})
	}
}
