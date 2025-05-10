package wordbreak

func wordBreak(s string, wordDict []string) bool {
    memoization := make(map[int]bool)


    var dp func(startIndex int) bool
    dp = func (startIndex int) bool {
        if startIndex == len(s) {
            return true
        }

        if val, ok := memoization[startIndex] ; ok {
            return val
        }

        for _, word := range wordDict {
            endIndex := startIndex + len(word)

            if endIndex > len(s) {
                continue
            }

            if s[startIndex:endIndex] == word {
                if dp(endIndex) {
                    memoization[startIndex] = true
                    return true
                }
            }
        }

        memoization[startIndex] = false
        return false
    }

    return dp(0)
}
