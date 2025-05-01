package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		for r > l && !isAlphaNum(string(s[l])) {
			l += 1
		}
		for r > l && !isAlphaNum(string(s[r])) {
			r -= 1
		}
		// !strings.EqualFold(string(s[l]), string(s[r]))
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}

		l++
		r--
	}

	return true

}

func isAlphaNum(s string) bool {
	return (('A' <= s[0] && s[0] <= 'Z') ||
		('a' <= s[0] && s[0] <= 'z') ||
		('0' <= s[0] && s[0] <= '9'))

}


