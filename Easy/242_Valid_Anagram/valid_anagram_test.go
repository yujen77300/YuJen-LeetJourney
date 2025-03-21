package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	test := []struct {
		name   string
		s      string
		t      string
		result bool
	}{
		{
			name:   "Example 1",
			s:      "anagram",
			t:      "nagaram",
			result: true,
		},
		{
			name:   "Example 2",
			s:      "rat",
			t:      "car",
			result: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.result {
				t.Errorf("isAnagram() = %v, want %v", result, tt.result)
			}
		})
	}
}
