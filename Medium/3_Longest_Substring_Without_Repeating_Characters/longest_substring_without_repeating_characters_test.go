package longestsubstringwithoutrepeatingcharacters

import "testing"


func TestLengthOfLongestSubstring(t *testing.T) {
	test := []struct {
		s        string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tt := range test {
		result := lengthOfLongestSubstring(tt.s)
		if result != tt.expected {
			t.Errorf("For input %s, expected %d but got %d", tt.s, tt.expected, result)
		}
	}
}
