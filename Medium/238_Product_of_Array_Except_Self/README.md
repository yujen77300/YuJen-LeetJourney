# [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)

## Solution 1: Brute Force with Two Loops
- **Time Complexity**: O(n²) - where n is the length of the array
- **Space Complexity**: O(n) - for the result array
- **Approach**:
   1. For each element in the array, compute the product of all other elements
   2. Use a nested loop to multiply all numbers except the current one
   3. This approach is straightforward but less efficient for large arrays

```go
func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    
    for i := 0; i < n; i++ {
        product := 1
        for j := 0; j < n; j++ {
            if j != i {
                product *= nums[j]
            }
        }
        result[i] = product
    }
    
    return result
}
```

## Solution 2: Prefix & Postfix Products
- **Time Complexity**: O(n) - where n is the length of the array
- **Space Complexity**: O(1) - excluding the output array
- **Approach**:
  1. Use two passes through the array to compute products
  2. First pass: Calculate prefix products (products of all elements to the left)
  3. Second pass: Calculate postfix products (products of all elements to the right)
  4. The result for each position is the product of its corresponding prefix and postfix values

```go
func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    prefix := 1
    for i := 0; i < n; i++ {
        res[i] = prefix
        prefix *= nums[i]
    }

    postfix := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}
```