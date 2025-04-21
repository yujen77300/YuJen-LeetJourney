# [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversalg/)

## Solution 1: BFS
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Each node is processed exactly once when it's dequeued
  - Adding and removing elements from the queue takes O(1) time
- **Space Complexity**: O(n) - for storing the queue and result
  - In the worst case (a complete binary tree's last level), the queue can hold up to n/2 nodes
  - The result list stores all n node values
- **Approach**:
  1. Use a queue data structure (FIFO - First In, First Out)
  2. Start by adding the root node to the queue
  3. Process nodes level by level, tracking the number of nodes at each level
  4. For each node, add its value to the current level list and enqueue its children
  5. After processing all nodes at a level, add the level list to the result
  6. Continue until the queue is empty


```go
func levelOrder(root *TreeNode) [][]int {


    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        queueLength := len(queue)
        level := []int{}


        for i := 0 ; i < queueLength ; i++ {
            node := queue[0]
            queue := queue[1:]
            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, level)
    }

    return result

    
}


```

## Solution 2: Two Pointers To Expand From Center
- **Time Complexity**: O(n²) - where n is the length of the string
  - We consider each position as a potential center: O(n)
  - For each center, we expand outward to find the longest palindrome: O(n)
  - Total: O(n) × O(n) = O(n²)
- **Space Complexity**: O(1) - only constant extra space used regardless of input size
- **Approach**:
  1. Iterate through each position in the string
  2. For each position, expand outward to find palindromes (checking both odd and even length cases)
  3. For odd-length palindromes, start with a single character as center
  4. For even-length palindromes, start with two adjacent characters
  5. Keep track of the longest palindrome found
  6. Return the longest palindromic substring

```go
func longestPalindrome(s string) string {
    result := ""
    maxLength := 0

    for i, _ := range s {
        l := i
        r := i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if (r-l+1) > maxLength{
                result = s[l:r+1]
                maxLength = r-l+1
            }

            l -= 1
            r += 1
        }

        l = i
        r = i + 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if (r-l+1) > maxLength{
                result = s[l:r+1]
                maxLength = r-l+1
            }

            l -= 1
            r += 1
        }
    } 

    return result
}
```