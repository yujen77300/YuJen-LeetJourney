# [Decode Ways](https://leetcode.com/problems/decode-ways/)

## Solution 1: Recursive Approach (Brute Force)
- **Time Complexity**: O(2^n) - where n is the length of the string
  - At each position, we make up to 2 recursive calls (decode 1 digit or 2 digits)
  - This creates a binary decision tree with potential depth n
  - Without memoization, many subproblems are recalculated
- **Space Complexity**: O(n) - for the recursion call stack
  - The maximum recursion depth equals n (when processing one character at a time)
- **Approach**:
  1. Use recursion with a decision tree approach
  2. At each step, try decoding the current digit as a single character (if valid)
  3. Also try decoding the current and next digits as a two-digit number (if valid)
  4. Base case: If we reach the end of the string, we've found a valid decoding (return 1)
  5. Base case: If the current digit is '0', it can't be decoded alone (return 0)
- **Key Insights**:
  - Each step presents a binary choice: decode 1 digit or 2 digits
  - A digit '0' can only be decoded as part of a two-digit number (10 or 20)
  - Valid two-digit decodings must be between 10-26 (corresponding to 'j' through 'z')
  - This recursive approach will exceed time limits for longer inputs due to redundant calculations

```go
func numDecodings(s string) int {
	var dfs func(index int) int
	dfs = func(index int) int {
		if index == len(s) {
			return 1
		}
		if s[index] == '0' {
			return 0
		}

    // Try taking single digit
		count := dfs(index + 1)

    // Try taking two digits if possible
		if index+1 < len(s) {
			if s[index] == '1' || (s[index] == '2' && s[index+1] >= '0' && s[index+1] <= '6') {
				count += dfs(index + 2)
			}
		}

		return count
	}

	return dfs(0)
}

```




## Solution 2: Memoization (Top-Down DP)
- **Time Complexity**: O(n) - where n is the length of the string
  - Each position in the string is processed at most once
  - The memoization prevents redundant calculations
- **Space Complexity**: O(n) - for memoization map and recursion stack
  - Store at most n results in the memoization map
  - The recursion stack can grow to depth n in worst casecomplexity
- **Approach**:
  1. Use the same recursive structure as Solution 1, but add memoization
  2. Store results for each index in a map to avoid recalculating
  3. Before computing a result, check if it's already in the memoization table
  4. After computing a result, store it in the table
  5. Return the memoized result for the starting index
  6. Return the total count
- **Key Insights**:
  - Memoization dramatically improves performance by eliminating redundant calculations

```go
func numDecodings(s string) int {
    dp := make(map[int]int)

    var dfs func(i int) int
    dfs = func(i int) int {
        if val, ok := dp[i]; ok {
            return val
        }

        if i == len(s) {
            return 1
        }

        if s[i] == '0' {
            return 0
        }

        res := dfs(i + 1)

        if i+1 < len(s) {
            if s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6') {
                res += dfs(i + 2)
            }
        }

        dp[i] = res
        return res
    }

    return dfs(0)
}

```


