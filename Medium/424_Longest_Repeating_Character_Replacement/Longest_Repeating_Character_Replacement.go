package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	countMap := make(map[byte]int)
	answer := 0
	maxFreq := 0
	rightIndex := 0

	// for leftIndex := 0; leftIndex < len(s); leftIndex++ {
	for leftIndex := range s {
		for rightIndex < len(s) && rightIndex-leftIndex-maxFreq <= k {
			countMap[s[rightIndex]] += 1
			maxFreq = max(maxFreq, countMap[s[rightIndex]])
			rightIndex += 1
		}

		// Update the length of the longest substring
		if rightIndex-leftIndex-maxFreq > k {
			// The new character is different from the current maxFreq character
			answer = max(answer, rightIndex-leftIndex-1)
		} else {
			// rightIndex has reached the rightmost position
			answer = max(answer, rightIndex-leftIndex)
		}

		countMap[s[leftIndex]] -= 1

	}

	return answer

}
