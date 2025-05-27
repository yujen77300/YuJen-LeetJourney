# [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

## Solution 1: Recursive Approach
- **Time Complexity**: O(2^n) - where n is the number of steps
  - Each step creates a binary tree of decisions (taking 1 or 2 steps)
  - The height of the recursion tree will be n, resulting in exponential complexity
- **Space Complexity**: O(n) - for the recursion stack
  - The maximum depth of recursion equals n (in the case of taking all 1-step moves)
- **Approach**:
  1. Use recursion to calculate all possible ways to reach the top
  2. Base case: If remaining steps is 0, we found a valid combination (return 1)
  3. Base case: If remaining steps is negative, this is an invalid path (return 0)
  4. For each step, we can either climb 1 step or 2 steps
  5. Sum the results of both choices


```go
func climbStairs(n int) int {

	var dfs func(int) int

	dfs = func(remaining int) int {
		if remaining < 0 {
			return 0
		}

		if remaining == 0 {
			return 1
		}

		return dfs(remaining-1) + dfs(remaining-2)

	}

	return dfs(n)

}
```


## Solution 2: Memoization Approach (Top-down DP)
- **Time Complexity**: O(n) - where n is the number of steps
  - Compute each subproblem exactly once
  - Each subproblem takes O(1) time to compute
- **Space Complexity**: O(n) - for the memoization table and recursion stack
  - Store at most n+1 states in the memo map
  - The recursion stack can grow up to depth n in the worst case
- **Approach**:
  1. Use a memoization table (map) to store results of subproblems
  2. Before calculating a subproblem, check if we've already solved it
  3. Use the same recursive structure as Solution 1, but avoid redundant calculations
  4. Base cases remain the same: 0 steps remaining returns 1, negative steps returns 0

```go

func climbStairs(n int) int {
    memo := make(map[int]int)

    var dfs func(int) int
    dfs = func(remaining int) int {
        if remaining < 0 {
            return 0
        }
        if remaining == 0 {
            return 1
        }

        if val, ok := memo[remaining]; ok {
            return val
        }

        memo[remaining] = dfs(remaining-1) + dfs(remaining-2)
        return memo[remaining]
    }

    return dfs(n)
}


```


## Solution 3: Bottom-up DP with Constant Space
- **Time Complexity**: O(n) - where n is the number of steps
  - Iterate through steps 1 to n exactly once
  - Each iteration performs constant-time operations
- **Space Complexity**: O(1) - constant extra space
  - Only three variables
- **Approach**:
  1. Recognize that this problem follows the Fibonacci sequence pattern
  2. Initialize first and second variables to 1 (base cases)
  3. Iterate from 1 to n-1, calculating the next number in the sequence
  4. At each step, the new value is the sum of the two previous values
  5. Return the final calculated value
- **Key Insights**:
  - The number of ways to climb n stairs equals the sum of ways to climb (n-1) and (n-2) stairs
  - Like Fibonacci sequence: F(n) = F(n-1) + F(n-2)
  - Only need to track the two most recent values, allowing constant space complexity


```go

func climbStairs(n int) int {
    first:= 1     // the number of ways to reach step 1 (dp[1])
    second := 1   // he number of ways to reach step 0 (dp[0])

    for i:=0 ; i < n-1 ; i++ {
        temp := first
        first = first + second
        second = temp
    }

    return first
}

```