# [House Robber](https://leetcode.com/problems/house-robber/)

## Solution 1: Top-down with Recursive (non-memoized)
- **Time Complexity**: O(2^n) - where n is the length of the array
  - Each house has two options: rob or skip
  - This creates a binary decision tree with depth n
  - Without memoization, many subproblems are recalculated multiple times
- **Space Complexity**: O(n) - for the recursion call stack
  - The maximum depth of recursion is n in the worst case
- **Approach**:
  1. Start from the last house (index n-1)
  2. At each house, decide whether to rob it or skip it
  3. If we rob the current house, we add its value to the maximum value we can get from houses [0...i-2]
  4. If we skip the current house, we get the maximum value from houses [0...i-1]
  5. Choose the maximum of these two options
  6. Base cases: For single house, return its value; for two houses, return the maximum of their values
- **Key Insights**:
  - This is a classic dynamic programming problem with overlapping subproblems
  - The recursive approach directly models the decision-making process
  - Without memoization, this solution is inefficient due to redundant calculations


```go
func rob(nums []int) int {
    n := len(nums)

    var helper func(i int)int
    helper = func (i int) int{
        if i == 0 {
            return nums[0]
        }
        if i == 1 {
            return max (nums[0],nums[1])
        }

        return max(helper(i-1), helper(i-2)+nums[i])

    }

    
    return helper(n-1)
}

```




## Solution 2: Top-Down with Recursive (memoized)
- **Time Complexity**: O(n) - where n is the length of the array
  - Each subproblem is solved only once and stored in the memoization map
  - With n possible indices, we have n subproblems
- **Space Complexity**: O(1) -  for the memoization map and recursion stack
  - The map stores at most n key-value pairs
  - The recursion stack can reach a maximum depth of n
- **Approach**:
  1. Use the same recursive approach as the non-memoized solution
  2. Store the results of subproblems in a map to avoid redundant calculations
  3. Before computing a subproblem, check if its result is already in the map
  4. Pre-compute and store base cases (for indices 0 and 1)
- **Key Insights**:
  - Memoization dramatically improves efficiency by eliminating redundant calculations


```go
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        return max(nums[0],nums[1])
    }

    memoMap := map[int]int{
        0 : nums[0],
        1 : max(nums[0],nums[1]),
    } 

    var helper func(i int)int
    helper = func (i int) int{
        if v, exists := memoMap[i]; exists{
            return v
        }else{
            memoMap[i] = max(helper(i-1), helper(i-2)+nums[i])
            return memoMap[i]
        }

    }

    
    return helper(n-1)
}
```




## Solution 3: Bottom-up
- **Time Complexity**: O(n²) - where n is the length of the string
  - Iterate through the array once, performing constant-time operations for each element
- **Space Complexity**: O(n) - for the dp array
  - Store n intermediate results in the dp array
- **Approach**:
  1. Create a dp array where dp[i] represents the maximum amount that can be robbed up to house i
  2. Initialize base cases: dp[0] = nums[0] and dp[1] = max(nums[0], nums[1])
  3. For each house from index 2 to n-1, calculate dp[i] as max(dp[i-1], dp[i-2] + nums[i])
  4. Return dp[n-1], which represents the maximum amount that can be robbed
- **Key Insights**:
  - Bottom-up approach builds the solution iteratively from smaller subproblems to larger ones
  - The dp array eliminates the need for recursion, making the solution more efficient
  - Each step only requires the results from the previous two steps, demonstrating optimal substructure


```go
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        return max(nums[0],nums[1])
    }


    dp := make([]int,n)
    dp[0] = nums[0]
    dp[1] = max(nums[0],nums[1])

    for i:=2 ; i < n ; i++ {
        dp[i] = max(dp[i-1],dp[i-2]+nums[i])
    }

    return dp[n-1]
    
}
```