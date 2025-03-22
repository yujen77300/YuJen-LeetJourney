package validparentheses

import "testing"

func TestIsValif(t *testing.T) {
	test := []struct {
		name   string
		s      string
		result bool
	}{
		{
			name:   "Example 1",
			s:      "()",
			result: true,
		},
		{
			name:   "Example 2",
			s:      "()[]{}",
			result: true,
		},
		{
			name:   "Example 3",
			s:      "(]",
			result: false,
		},
		{
			name:   "Example 4",
			s:      "([])",
			result: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.s)
			if result != tt.result {
				t.Errorf("isValid() = %v, want %v", result, tt.result)
			}
		})
	}
}
