# [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

## Solution 1: Recursive Approach
- **Time Complexity**: O(2^n) - where n is the input number
  - Each function call branches into two recursive calls, creating exponential growth
  - The recursion tree has approximately 2^n nodes
- **Space Complexity**: O(n) - for the recursion stack
  - The maximum depth of recursion equals n (worst case path)
- **Approach**:
  1. Use the mathematical definition of Fibonacci directly
  2. Base case: F(0) = 0, F(1) = 1
  3. Recursive case: F(n) = F(n-1) + F(n-2)
  4. Let recursion handle the computation tree
- **Key Insights**:
  - Same subproblems are solved multiple times


```go
func fib(n int) int {
    if n < 2 {
        return n
    }

    return fib(n-1) + fib(n-2)
}
```


## Solution 2: Memoization Approach (Top-down DP)
- **Time Complexity**: O(n) - where n is the input number
  - Each subproblem from 0 to n is computed exactly once
  - Each computation takes O(1) time after memoization
- **Space Complexity**: O(n) - for the memoization table and recursion stack
  - Memoization table stores at most n+1 values
  - Recursion stack depth can reach n in the worst case
- **Approach**:
  1. Use a memoization table (map) to cache computed results
  2. Before computing F(n), check if it's already in the cache
  3. If not cached, compute recursively and store the result
  4. Return the cached or newly computed value
- **Key Insights**:
  - Top-down approach: solve larger problems by breaking into smaller ones
  - Eliminates redundant calculations from naive recursion


```go
func fib(n int) int {
	memo := make(map[int]int)

	var f func(int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}

		if v, exist := memo[n]; exist {
			return v
		}

		memo[n] = f(n-1) + f(n-2)
		return memo[n]
	}

	return f(n)

}


```


## Solution 3: Tabulation Approach (Bottom-up DP)
- **Time Complexity**: O(n) - where n is the input number
  - Single loop from 0 to n, each iteration takes constant time
- **Space Complexity**: O(n) - for the DP table
  - Array of size n+1 to store all Fibonacci numbers from 0 to n
- **Approach**:
  1. Create a table to store Fibonacci numbers from 0 to n
  2. Fill base cases: table[0] = 0, table[1] = 1
  3. Iteratively fill the table using the recurrence relation
  4. Each entry depends only on the two previous entries
  5. Return the final computed value
- **Key Insights**:
  - Bottom-up approach: build solutions from smallest to largest subproblems
  - Eliminates recursion overhead completely
  - All subproblems are solved in optimal order



```go

func fib(n int) int {
    table := make([]int, n+1)

    for i := 0 ; i <= n ; i++ {
        if i == 0 {
            table[i] = 0
            continue
        }

        if i == 1 {
            table[i] = 1
            continue
        }


        table[i] = table[i-1] + table[i-2]
    }

    return table[n]


}

```



## Solution 4: Space-Optimized Tabulation (Bottom-up DP)
- **Time Complexity**: O(n) - where n is the input number
  - Single loop from 2 to n, each iteration takes constant time
- **Space Complexity**: O(1) - constant extra space
  - Only two variables needed to track the previous two Fibonacci numbers
- **Approach**:
  1. Recognize we only need the last two values to compute the next one
  2. Use two variables (prev, cur) instead of an entire array
  3. Iteratively update these variables using simultaneous assignment
  4. At each step, shift values: prev becomes cur, cur becomes prev+cur
- **Key Insights**:
  - Space optimization: only store what's necessary for the next computation

```go

func fib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    prev, cur := 0, 1
    for i := 2; i <= n; i++ {
        prev, cur = cur, prev+cur
    }
    return cur
}

```