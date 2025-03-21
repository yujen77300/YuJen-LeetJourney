package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for _, v := range s {
		charCount[v]++
	}

	for _, v := range t {
		charCount[v]--
		if charCount[v] < 0 {
			return false
		}
	}

	return true
}
