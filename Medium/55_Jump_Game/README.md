# [Jump Game](https://leetcode.com/problems/jump-game/)

## Solution 1: Recursive
- **Time Complexity**: O(k^n) - `k = max(nums[i])`,  `n = len(nums)`
  - At each position, we can make up to `nums[i]` different jumps
  - The maximum number of choices at any position is the maximum value in the array `k = max(nums[i])`
  - In the worst case, we might need to explore most of these choices to a depth of `n`
- **Space Complexity**: O(n) - for the recursion call stack
  - The maximum depth of recursion is n in the worst case
- **Approach**:
  1. Use a recursive function to check if we can reach the end
  2. For each position, try all possible jumps (from 1 to nums[i])
  3. If any jump leads to a position that can reach the end, return true
  4. Return false if no valid jumps are found.

```go
func canJump(nums []int) bool {
	n := len(nums)

	var canReach func(i int) bool
	canReach = func(i int) bool {
		if i == n-1 {
			return true
		}
		for jump := 1; jump <= nums[i]; jump++ {
			if i+jump < n && canReach(i+jump) {
				return true
			}
		}
		return false
	}

	return canReach(0)
}

```




## Solution 2: Greedy Approach
- **Time Complexity**: O(n) - where n is the length of the array
  - Iterate through the array once from right to left
  - Each element is processed in constant time
- **Space Complexity**: O(1) - constant extra space
  - Only a single variable (goal) regardless of input size
- **Approach**:
  1. Start with the goal at the last position
  2. Iterate backwards from the second-to-last position
  3. At each position, check if a jump can reach the current goal
  4. If reachable, update the goal to the current position
  5. After iteration, check if the goal has reached the start (index 0)
- **Key Insights**:
  - Only concerned with reaching the end, not finding the optimal path

```go
func canJump(nums []int) bool {
    goal := len(nums) - 1

    for i := goal - 1 ; i >= 0 ; i -- {
        maxJump := nums[i]
        if i + maxJump >= goal {
            goal = i
        }
    }

    return goal == 0

}

```


