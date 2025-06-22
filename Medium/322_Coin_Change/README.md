# [Coin Change](https://leetcode.com/problems/coin-change/)

## Solution 1: Dynamic Programming
- **Time Complexity**: O(amount × n) - where amount is the target amount and n is the number of coin denominations
  - For each amount from 1 to target amount (O(amount)), we try each coin denomination (O(n))
  - This results in O(amount × n) operations
- **Space Complexity**: O(amount) - linear extra space
  - Only a dp array of size amount+1 to store the minimum number of coins needed for each amount
- **Approach**:
  1. Create a dp array where dp[i] represents the minimum number of coins needed to make amount i
  2. Initialize dp[0] = 0 (base case: 0 coins needed to make amount 0)
  3. Initialize all other dp values to amount+1 (representing unreachable)
  4. For each amount i from 1 to target amount:
     - For each coin denomination:
       - If the current amount i is at least the coin value, we can use this coin
       - Update dp[i] to the minimum of its current value or dp[i-coin] + 1
  5. Return dp[amount] if it's less than or equal to amount (valid solution), otherwise return -1
- **Key Insights**:
  - Build solutions for smaller amounts first, then use them to solve for larger amounts
  - Each dp[i] represents the optimal solution for amount i
  - Using amount+1 as "unreachable" is a common technique that simplifies the code
  - Take coins = [1,2,5], amount = 6, to draw a dp table

| Amount / Coin | 1           | 2                             | 3                         | 4            | 5                         | 6                         |
| ------------- | ----------- | ----------------------------- | ------------------------- | ------------ | ------------------------- | ------------------------- |
| **1**         | 1 coin of 1 | 2 coins of 1                  | 3 coins of 1              | 4 coins of 1 | 5 coins of 1              | 6 coins of 1              |
| **2**         | 1 coin of 1 | `min(1×2, 2×1) = 1 coin of 2` | 1×2 + 1×1 → total 2 steps | 2×2          | 2×2 + 1×1 → total 3 steps | 3×2                       |
| **5**         | 1 coin of 1 | `min(1×2, 2×1) = 1 coin of 2` | 1×2 + 1×1 → total 2 steps | 2×2          | 1 coin of 5               | 1×5 + 1×1 → total 2 steps |



```go
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// Initialize (dp[0]=0, others set to a large number)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
```



