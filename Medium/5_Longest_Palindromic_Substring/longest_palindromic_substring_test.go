package longestpalindromicsubstring

import "testing"


func TestLongestPalindrome(t *testing.T) {
	test := []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, tt := range test {
		result := longestPalindrome(tt.s)
		if result != tt.expected {
			t.Errorf("Expected %s, but got %s", tt.expected, result)
		}
	}
}
